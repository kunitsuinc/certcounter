package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/kunitsuinc/rec.go"
	contextz "github.com/kunitsuinc/util.go/context"
	"github.com/kunitsuinc/util.go/must"

	"github.com/kunitsuinc/certcounter/pkg/config"
	"github.com/kunitsuinc/certcounter/pkg/consts"
	"github.com/kunitsuinc/certcounter/pkg/entrypoint"
	"github.com/kunitsuinc/certcounter/pkg/errors"
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

//nolint:cyclop
func Main(ctx context.Context, l *rec.Logger) error {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
	ctx = contextz.WithSignalChannel(ctx, signalChan)

	_, errConfigLoad := config.Load(l)

	// NOTE: If -h option, return nil
	if errors.Is(errConfigLoad, flag.ErrHelp) {
		return nil
	}

	// NOTE: display version info
	l.Info("version info", rec.Object("version", map[string]string{"Version": config.Version(), "Revision": config.Revision(), "Branch": config.Branch(), "Timestamp": config.Timestamp(), "GoVersion": config.GoVersion()}))

	// NOTE: If -v option, return nil
	if config.SubcommandVersion() {
		return nil
	}

	l.With(rec.String("command", os.Args[0]), rec.Strings("args", os.Args[1:]), rec.Int("pid", os.Getpid())).F().Infof("main: 🔆 start %s command=%s args=%v pid=%d", consts.AppName, os.Args[0], os.Args[1:], os.Getpid())
	defer func() {
		l.With(rec.String("command", os.Args[0]), rec.Strings("args", os.Args[1:]), rec.Int("pid", os.Getpid())).F().Infof("main: 💤 shutdown %s command=%s args=%v pid=%d", consts.AppName, os.Args[0], os.Args[1:], os.Getpid())
	}()

	// NOTE: If err != nil, panic(err)
	must.Must(errConfigLoad)

	if err := entrypoint.CertCounter(ctx, l); err != nil {
		return errors.Errorf("entrypoint.CertCounter: %w", err)
	}

	return nil
}
