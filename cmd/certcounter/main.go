package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/kunitsuinc/certcounter/pkg/config"
	"github.com/kunitsuinc/certcounter/pkg/entrypoint"
	"github.com/kunitsuinc/rec.go"
	"github.com/kunitsuinc/util.go/must"
)

func main() {
	l := rec.Must(rec.New(os.Stdout))
	rec.ReplaceDefaultLogger(l)

	shutdownChan := make(chan os.Signal, 1)
	signal.Notify(shutdownChan, syscall.SIGINT, syscall.SIGTERM)

	ctx := context.Background()
	config.Load()

	if config.SubcommandVersion() {
		fmt.Fprintf(os.Stdout, "%s\n%s\n%s\n%s\n", config.Version(), config.Revision(), config.Branch(), config.Timestamp())
		return
	}

	must.Must(config.Check())

	if err := entrypoint.CertCounter(ctx); err != nil {
		l.F().Errorf("entrypoint.CertCounter: %v", err)
		os.Exit(1)
	}
}
