package entrypoint

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/kunitsuinc/certcounter/pkg/config"
	"github.com/kunitsuinc/certcounter/pkg/consts"
	"github.com/kunitsuinc/certcounter/pkg/controller/router"
	"github.com/kunitsuinc/certcounter/pkg/errors"
	"github.com/kunitsuinc/certcounter/pkg/traces"
	"github.com/kunitsuinc/grpcutil.go/grpcz"
	"github.com/kunitsuinc/rec.go"
	"github.com/kunitsuinc/util.go/must"
	"golang.org/x/net/http2"
)

func CertCounter(ctx context.Context, l *rec.Logger) (serve func(errChan chan<- error), shutdown func() error) {
	gcpProjectID := config.GoogleCloudProject()
	awsProfile := config.AWSProfile()
	shutdownTimeout := config.ShutdownTimeout()

	_ = gcpProjectID
	_ = awsProfile

	// nolint: contextcheck
	shutdownTracerProvider := traces.InitTracerProvider(traces.NewExporter(l), traces.NewResource(consts.AppName, config.Version()), l)

	address := fmt.Sprintf("%s:%d", config.Addr(), config.Port())

	grpcGatewayRouter := must.One(router.NewGRPCGatewayRouter(ctx, address, l))

	mux := http.NewServeMux()
	mux.Handle("/", grpcGatewayRouter)

	grpcServer := router.NewGRPCServer(l)

	server := &http.Server{
		Addr:              address,
		Handler:           grpcz.GRPCHandler(grpcServer, mux, &http2.Server{}),
		ReadHeaderTimeout: 10 * time.Second,
	}

	serve = func(errChan chan<- error) {
		l.F().Infof("🔊 start gRPC Server with gRPC-Gateway: %s", address)
		defer l.F().Infof("🔇 shutdown gRPC Server and gRPC-Gateway: %s", address)

		if err := server.ListenAndServe(); err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				l.Info(err.Error())
				errChan <- nil
				return
			}

			errChan <- errors.Errorf("http.Serve: %w", err)
			return
		}

		errChan <- nil
	}

	shutdown = func() error {
		grpcServer.GracefulStop()

		shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
		defer cancel()
		if err := server.Shutdown(shutdownCtx); err != nil {
			return errors.Errorf("server.Shutdown: %w", err)
		}

		shutdownTracerProvider()

		return nil
	}

	return serve, shutdown
}
