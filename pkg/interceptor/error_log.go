package interceptor

import (
	"context"

	statusz "github.com/kunitsuinc/grpcutil.go/grpc/status"
	"github.com/kunitsuinc/rec.go"
	"google.golang.org/grpc"

	"github.com/kunitsuinc/certcounter/pkg/errors"
)

func ErrorLogInterceptor(original *rec.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		resp, err := handler(ctx, req)
		if err != nil {
			l := rec.ContextLogger(ctx)
			sz := &statusz.Status{}
			if errors.As(err, &sz) {
				return nil, l.E().Error(errors.Errorf("rpc error: code = %s desc = %w", sz.GRPCStatus(), sz)).Err()
			}
			return nil, l.E().Error(err).Err()
		}

		return resp, nil
	}
}
