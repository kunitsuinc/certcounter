package router

import (
	"context"
	"net/http"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"

	// grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator".
	gw_runtime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/kunitsuinc/grpcutil.go/grpc-ecosystem/go-grpc-middleware/errorhandler"
	gw_runtimez "github.com/kunitsuinc/grpcutil.go/grpc-ecosystem/grpc-gateway/v2/runtime"
	statusz "github.com/kunitsuinc/grpcutil.go/grpc/status"
	"github.com/kunitsuinc/rec.go"
	"github.com/kunitsuinc/util.go/net/httpz"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/keepalive"

	v1 "github.com/kunitsuinc/certcounter/generated/go/certcounter/v1"
	"github.com/kunitsuinc/certcounter/pkg/config"
	"github.com/kunitsuinc/certcounter/pkg/consts"
	"github.com/kunitsuinc/certcounter/pkg/controller"
	"github.com/kunitsuinc/certcounter/pkg/errors"
	"github.com/kunitsuinc/certcounter/pkg/interceptor"
	"github.com/kunitsuinc/certcounter/pkg/middleware"
)

func registerGRPCServer(grpcServer *grpc.Server) *grpc.Server {
	// NOTE: register service (1/2)
	v1.RegisterTestAPIServiceServer(grpcServer, &controller.TestAPIController{})

	return grpcServer
}

func registerHandlers(ctx context.Context, mux *gw_runtime.ServeMux, conn *grpc.ClientConn) error {
	if err := gw_runtimez.RegisterHandlers(ctx, mux, conn,
		// NOTE: register service (2/2)
		v1.RegisterTestAPIServiceHandler,
	); err != nil {
		return errors.Errorf("gw_runtimez.RegisterHandlers: %w", err)
	}

	return nil
}

func NewGRPCServer(l *rec.Logger) *grpc.Server {
	s := registerGRPCServer(grpc.NewServer(
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				// NOTE: https://github.com/grpc-ecosystem/go-grpc-middleware
				otelgrpc.UnaryServerInterceptor(), // NOTE: OpenTelemetry for gRPC Gateway -> gRPC Server
				grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
				interceptor.LoggerInterceptor(l),
				interceptor.AccessLogInterceptor(),
				errorhandler.UnaryServerInterceptor(func(ctx context.Context, info *grpc.UnaryServerInfo, err error) error {
					l := rec.ContextLogger(ctx)
					sz := &statusz.Status{}
					if errors.As(err, &sz) {
						return l.With(rec.Uint32("grpc.code", uint32(sz.GRPCStatus().Code())), rec.String("grpc.message", sz.GRPCStatus().Message())).E().Error(sz).Err()
					}
					return l.E().Error(err).Err()
				}),
				// grpc_validator.UnaryServerInterceptor(),
			),
		),
	))

	// NOTE: healthcheck
	healthSrv := health.NewServer()
	grpc_health_v1.RegisterHealthServer(s, healthSrv)
	healthSrv.SetServingStatus(consts.AppName, grpc_health_v1.HealthCheckResponse_SERVING)

	return s
}

// NewRouter is
// TODO
//
// cf. https://github.com/grpc-ecosystem/grpc-gateway
func NewRouter(ctx context.Context, grpcServerEndpoint string, l *rec.Logger) (http.Handler, error) {
	conn, err := grpc.Dial(
		grpcServerEndpoint,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(
			grpc_middleware.ChainUnaryClient(
				otelgrpc.UnaryClientInterceptor(), // NOTE: OpenTelemetry for gRPC Gateway -> gRPC Server
			),
		),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:                1 * time.Minute,
			Timeout:             1 * time.Minute,
			PermitWithoutStream: true,
		}),
	)
	if err != nil {
		return nil, errors.Errorf("grpc.Dial: %w", err)
	}

	mux := gw_runtime.NewServeMux(
		gw_runtime.WithHealthzEndpoint(grpc_health_v1.NewHealthClient(conn)),
	)
	if err := registerHandlers(ctx, mux, conn); err != nil {
		return nil, err
	}

	middlewares := httpz.
		Middlewares(
			httpz.NewResponseWriterBufferHandler(func(rwb *httpz.ResponseWriterBuffer, r *http.Request) {
				l := rec.ContextLogger(r.Context())
				l.With(rec.Int("statusCode", rwb.StatusCode()), rec.Int64("contentLength", int64(rwb.Buffer.Len()))).F().Infof("access: %d %s (Content-Length:%d) <- %s %s (Content-Length:%d)", rwb.StatusCode(), http.StatusText(rwb.StatusCode()), rwb.Buffer.Len(), r.Method, r.URL.Path, r.ContentLength)
			}).Middleware,
		).
		Middlewares(middleware.ContextLoggerRequestMiddleware(l)).
		Middlewares(httpz.NewXRealIPHandler(config.SetRealIPFrom(), httpz.HeaderXForwardedFor, true).Middleware)

	// NOTE: OpenTelemetry for client -> gRPC Gateway
	otelHandler := otelhttp.NewHandler(
		middlewares(mux),
		"gRPC-Gateway",
		otelhttp.WithTracerProvider(otel.GetTracerProvider()),
		otelhttp.WithPropagators(otel.GetTextMapPropagator()),
		otelhttp.WithSpanNameFormatter(func(operation string, r *http.Request) string { return operation + r.URL.Path }),
	)

	return otelHandler, nil
}
