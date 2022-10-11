package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kunitsuinc/certcounter/pkg/config"
	"github.com/kunitsuinc/certcounter/pkg/consts"
	"github.com/kunitsuinc/certcounter/pkg/entrypoint"
	"github.com/kunitsuinc/certcounter/pkg/errors"
	"github.com/kunitsuinc/rec.go"
	"github.com/kunitsuinc/util.go/contextz"
	"github.com/kunitsuinc/util.go/must"
)

func main() {
	l := rec.Must(rec.New(os.Stdout))
	rec.ReplaceDefaultLogger(l)

	ctx := rec.ContextWithLogger(context.Background(), l)

	if err := Main(ctx, l); err != nil {
		l.With(rec.Error(err)).F().Errorf("certcounter: %v", err)
		os.Exit(1)
	}
}

func Main(ctx context.Context, l *rec.Logger) error {
	shutdownChan := make(chan os.Signal, 1)
	signal.Notify(shutdownChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
	ctx = contextz.WithSignalChannel(ctx, shutdownChan)

	rollback, err := config.Load(l)
	defer rollback()

	// NOTE: If -h option, return nil
	if errors.Is(err, flag.ErrHelp) {
		return nil
	}

	// NOTE: display version info
	l.Info("version info", rec.Object("version", map[string]string{"Version": config.Version(), "Revision": config.Revision(), "Branch": config.Branch(), "Timestamp": config.Timestamp(), "GoVersion": config.GoVersion()}))

	// NOTE: If -v option, return nil
	if config.SubcommandVersion() {
		return nil
	}

	// NOTE: If err != nil, panic(err)
	must.Must(err)

	l.F().Infof("main: 🔆 start %s (pid:%d)", consts.AppName, os.Getpid())
	defer func() { l.F().Infof("main: 💤 shutdown %s (pid:%d)", consts.AppName, os.Getpid()) }()

	shutdown, errChanCertCounter := entrypoint.CertCounter(ctx, l)
	defer func() {
		shutdown()
		time.Sleep(1 * time.Millisecond) // NOTE: wait shutdown log
	}()

SignalLoop:
	for {
		select {
		case <-ctx.Done():
			l.E().Info(ctx.Err())
			break SignalLoop
		case sig := <-shutdownChan:
			l.F().Infof("main: catch the signal: %s", sig)
			if sig == syscall.SIGHUP {
				l.Info("main: reload config")
				rollback, err := config.Load(l)
				if err != nil {
					l.Info("main: failed to load config. rollback")
					rollback()
				}
				continue
			}
			break SignalLoop
		case err := <-errChanCertCounter:
			if err != nil {
				return errors.Errorf("main: entrypoint.CertCounter: %w", err)
			}
		}
	}

	return nil
}
