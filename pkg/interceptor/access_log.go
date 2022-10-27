package interceptor

import (
	"context"
	"strconv"

	"github.com/kunitsuinc/rec.go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func AccessLogInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		l := rec.ContextLogger(ctx)

		panicked := true

		defer func() {
			var code codes.Code
			if err != nil {
				if s, ok := status.FromError(err); ok {
					code = s.Code()
				}
			}
			if panicked {
				code = codes.Internal
			}

			// NOTE: 無名関数で囲わないと err == nil や panicked == true をキャプチャしてしまう
			success := err == nil && !panicked
			l.Info("access: "+code.String()+" code="+strconv.FormatUint(uint64(code), 10)+" success="+strconv.FormatBool(success)+" grpc.fullMethod="+info.FullMethod, rec.Uint32("grpc.code", uint32(code)), rec.Bool("success", success))
		}()

		resp, err = handler(ctx, req)
		panicked = false
		if err != nil {
			return nil, err
		}

		return resp, nil
	}
}
