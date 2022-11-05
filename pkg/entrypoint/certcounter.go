package entrypoint

import (
	"context"
	"net"
	"net/http"
	"os"
	"strconv"
	"syscall"

	grpcz "github.com/kunitsuinc/grpcutil.go/grpc"
	"github.com/kunitsuinc/rec.go"
	"github.com/kunitsuinc/util.go/must"

	"github.com/kunitsuinc/certcounter/pkg/config"
	"github.com/kunitsuinc/certcounter/pkg/controller"
	"github.com/kunitsuinc/certcounter/pkg/errors"
	"github.com/kunitsuinc/certcounter/pkg/traces"
)

func CertCounter(ctx context.Context, l *rec.Logger) error {
	gcpProjectID := config.GoogleCloudProject()
	awsProfile := config.AWSProfile()
	shutdownTimeout := config.ShutdownTimeout()

	_ = gcpProjectID
	_ = awsProfile

	//nolint:contextcheck
	shutdownTracerProvider := traces.InitTracerProvider(l)
	defer shutdownTracerProvider()

	address := net.JoinHostPort(config.Addr(), strconv.Itoa(config.Port()))

	mux := http.NewServeMux()
	mux.Handle("/", must.One(controller.NewRouter(ctx, address, l)))

	l.F().Infof("🔊 start gRPC Server with gRPC-Gateway: %s", address)
	defer l.F().Infof("🔇 shutdown gRPC Server with gRPC-Gateway: %s", address)

	if err := grpcz.ServeGRPC(
		ctx,
		must.One(net.Listen("tcp", address)),
		controller.NewGRPCServer(l),
		mux,
		grpcz.WithContinueSignalHandler(func(sig os.Signal) bool {
			if sig == syscall.SIGHUP {
				l.Info("main: reload config")
				rollback, err := config.Load(l)
				if err != nil {
					l.Warning("main: failed to load config. rollback")
					rollback()
				}
				return true
			}
			return false
		}),
		grpcz.WithShutdownTimeout(shutdownTimeout),
		grpcz.WithShutdownErrorHandler(func(err error) {
			l.With(rec.Error(err), rec.ErrorStacktrace(err)).F().Errorf("main: %v", err)
		}),
	); err != nil {
		return errors.Errorf("grpcz.ServeGRPC: %w", err)
	}

	return nil
}
