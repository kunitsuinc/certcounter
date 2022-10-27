package interceptor

import (
	"context"
	"path"

	"github.com/kunitsuinc/rec.go"
	"google.golang.org/grpc"
)

func LoggerInterceptor(original *rec.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		ctxWithLogger := rec.ContextWithLogger(ctx,
			original.With(
				rec.String("grpc.service", path.Dir(info.FullMethod)[1:]),
				rec.String("grpc.method", path.Base(info.FullMethod)),
				rec.String("grpc.fullMethod", info.FullMethod),
			),
		)

		resp, err := handler(ctxWithLogger, req)
		if err != nil {
			return nil, err
		}

		return resp, nil
	}
}
