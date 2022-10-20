package controller

import (
	"context"

	v1 "github.com/kunitsuinc/certcounter/generated/go/certcounter/v1"
	"github.com/kunitsuinc/certcounter/pkg/traces"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/emptypb"
)

type HealthController struct {
	v1.UnimplementedHealthServiceServer
}

func (*HealthController) Check(ctx context.Context, _ *emptypb.Empty) (*v1.HealthServiceCheckResponse, error) {
	_, span := traces.Start(ctx, "Check")
	defer span.End()

	c := codes.OK

	return &v1.HealthServiceCheckResponse{
		Code:    uint32(c),
		Message: c.String(),
	}, nil
}
