package interceptor

import (
	"context"
	"path"

	"github.com/kunitsuinc/rec.go"
	"google.golang.org/grpc"
)

func AccessLogInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		l := rec.ContextLogger(ctx)
		l = l.With(
			rec.String("grpc.service", path.Dir(info.FullMethod)[1:]),
			rec.String("grpc.method", path.Base(info.FullMethod)),
			rec.String("grpc.fullMethod", info.FullMethod),
		)

		panicked := true

		defer func() {
			// NOTE: 無名関数で囲わないと err == nil や panicked == true をキャプチャしてしまう
			success := err == nil && !panicked
			l.With(rec.Bool("success", success)).F().Infof("access: grpc.method=%s success=%t", info.FullMethod, success)
		}()

		resp, err = handler(ctx, req)
		panicked = false
		if err != nil {
			return nil, err
		}

		return resp, nil
	}
}
