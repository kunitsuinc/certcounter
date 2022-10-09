package entrypoint

import (
	"context"
	"fmt"
	"net/http"

	"github.com/kunitsuinc/certcounter/pkg/config"
	"github.com/kunitsuinc/certcounter/pkg/consts"
	"github.com/kunitsuinc/certcounter/pkg/controller/router"
	"github.com/kunitsuinc/certcounter/pkg/errors"
	"github.com/kunitsuinc/certcounter/pkg/traces"
	"github.com/kunitsuinc/grpcutil.go/grpcz"
	"github.com/kunitsuinc/rec.go"
	"golang.org/x/net/http2"
)

func CertCounter(ctx context.Context, l *rec.Logger) (shutdown func(), errChan <-chan error) {
	gcpProjectID := config.GoogleCloudProject()
	awsProfile := config.AWSProfile()
	shutdownTimeout := config.ShutdownTimeout()

	_ = gcpProjectID
	_ = awsProfile

	// nolint: contextcheck
	shutdownTracerProvider := traces.InitTracerProvider(traces.NewExporter(l), traces.NewResource(consts.AppName, config.Version()), l)

	errCh := make(chan error, 1)

	address := fmt.Sprintf("%s:%d", config.Addr(), config.Port())

	grpcGatewayRouter, err := router.NewGRPCGatewayRouter(ctx, address, l)
	if err != nil {
		errCh <- errors.Errorf("router.NewGRPCGatewayRouter: %w", err)
		return func() {}, errCh
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcGatewayRouter)

	grpcServer := router.NewGRPCServer(l)

	handler := grpcz.GRPCHandler(grpcServer, mux, &http2.Server{})

	server := &http.Server{Addr: address, Handler: handler}

	go func() {
		l.F().Infof("🔊 start gRPC Server with gRPC-Gateway: %s", address)
		defer l.F().Infof("🔇 shutdown gRPC Server and gRPC-Gateway: %s", address)

		if err := server.ListenAndServe(); err != nil {
			errCh <- errors.Errorf("http.Serve: %w", err)
			return
		}
		errCh <- nil
		return // nolint: gosimple
	}()

	shutdown = func() {
		grpcServer.GracefulStop()

		shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
		defer cancel()
		if err := server.Shutdown(shutdownCtx); err != nil {
			l.E().Error(errors.Errorf("server.Shutdown: %w", err))
			return
		}

		shutdownTracerProvider()
	}

	return shutdown, errCh
}
