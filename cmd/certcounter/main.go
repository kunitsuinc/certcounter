package main

import (
	"context"
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
	signal.Notify(shutdownChan, syscall.SIGINT, syscall.SIGTERM)
	ctx = contextz.WithSignalChannel(ctx, shutdownChan)

	config.Load(l)

	l.Info("version info",
		rec.Object("version", map[string]string{
			"Version":   config.Version(),
			"Revision":  config.Revision(),
			"Branch":    config.Branch(),
			"Timestamp": config.Timestamp(),
			"GoVersion": config.GoVersion(),
		}),
	)

	if config.SubcommandVersion() {
		return nil
	}

	l.F().Infof("main: 🔆 start %s", consts.AppName)
	defer l.F().Infof("main: 💤 shutdown %s", consts.AppName)

	must.Must(config.Check())

	shutdown, errCh := entrypoint.CertCounter(ctx, l)

	select {
	case <-ctx.Done():
		l.E().Info(ctx.Err())
	case sig := <-shutdownChan:
		l.F().Infof("catch the signal: %s", sig)
	case err := <-errCh:
		if err != nil {
			return errors.Errorf("entrypoint.StartGRPCGatewayAsync: %w", err)
		}
	}

	shutdown()

	time.Sleep(1 * time.Millisecond) // NOTE: wait shutdown log

	return nil
}
