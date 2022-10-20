package router

import (
	"context"
	"net"
	"net/http"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	gw_runtime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	v1 "github.com/kunitsuinc/certcounter/generated/go/certcounter/v1"
	"github.com/kunitsuinc/certcounter/pkg/controller"
	"github.com/kunitsuinc/certcounter/pkg/errors"
	"github.com/kunitsuinc/certcounter/pkg/interceptor"
	"github.com/kunitsuinc/certcounter/pkg/middleware"
	gw_runtimez "github.com/kunitsuinc/grpcutil.go/grpc-gateway/v2/runtimez"
	"github.com/kunitsuinc/rec.go"
	"github.com/kunitsuinc/util.go/net/httpz"
	"github.com/kunitsuinc/util.go/netz"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
)

func NewGRPCServer(l *rec.Logger) *grpc.Server {
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				// NOTE: https://github.com/grpc-ecosystem/go-grpc-middleware
				otelgrpc.UnaryServerInterceptor(), // NOTE: OpenTelemetry for gRPC Gateway -> gRPC Server
				grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
				interceptor.LoggerInterceptor(l),
				interceptor.AccessLogInterceptor(),
				interceptor.ErrorLogInterceptor(l),
				grpc_validator.UnaryServerInterceptor(),
			),
		),
	)

	// NOTE: register servers
	v1.RegisterTestAPIServiceServer(grpcServer, &controller.TestAPIController{})
	v1.RegisterHealthServiceServer(grpcServer, &controller.HealthController{})

	return grpcServer
}

// NewGRPCGatewayRouter TODO
// cf. https://github.com/grpc-ecosystem/grpc-gateway
func NewGRPCGatewayRouter(ctx context.Context, grpcServerEndpoint string, l *rec.Logger) (http.Handler, error) {
	mux := gw_runtime.NewServeMux()
	opts := []grpc.DialOption{
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
	}

	if err := gw_runtimez.RegisterHandlersFromEndpoint(ctx, mux, grpcServerEndpoint, opts,
		// NOTE: register handlers
		v1.RegisterTestAPIServiceHandlerFromEndpoint,
		v1.RegisterHealthServiceHandlerFromEndpoint,
	); err != nil {
		return nil, errors.Errorf("gw_runtimez.RegisterHandlersFromEndpoint: %w", err)
	}

	middlewares := httpz.
		Middlewares(httpz.NewResponseWriterBufferHandler(func(rwb *httpz.ResponseWriterBuffer, r *http.Request) {
			l := rec.ContextLogger(r.Context())
			l.With(rec.Int("statusCode", rwb.StatusCode), rec.Int64("contentLength", int64(rwb.Buffer.Len()))).F().Infof("access: %d %s (Content-Length:%d) <- %s %s (Content-Length:%d)", rwb.StatusCode, http.StatusText(rwb.StatusCode), rwb.Buffer.Len(), r.Method, r.URL.Path, r.ContentLength)
		}).Middleware).
		Middlewares(middleware.ContextLoggerRequestMiddleware(l)).
		Middlewares(httpz.NewXRealIPHandler([]*net.IPNet{netz.PrivateIPAddressClassA}, "", true).Middleware)

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
