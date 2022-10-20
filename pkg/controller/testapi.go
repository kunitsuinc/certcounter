package controller

import (
	"context"

	v1 "github.com/kunitsuinc/certcounter/generated/go/certcounter/v1"
	"github.com/kunitsuinc/certcounter/pkg/errors"
	"github.com/kunitsuinc/certcounter/pkg/traces"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type TestAPIController struct {
	v1.UnimplementedTestAPIServiceServer
}

func (*TestAPIController) Echo(ctx context.Context, request *v1.TestAPIServiceEchoRequest) (*v1.TestAPIServiceEchoResponse, error) {
	_, span := traces.Start(ctx, "Echo")
	defer span.End()

	return &v1.TestAPIServiceEchoResponse{
		Message: request.GetMessage(),
	}, nil
}

func (*TestAPIController) EchoError(ctx context.Context, request *v1.TestAPIServiceEchoErrorRequest) (*emptypb.Empty, error) {
	_, span := traces.Start(ctx, "EchoError")
	defer span.End()

	s := status.New(codes.Code(request.GetCode()), request.GetMessage())

	d, err := s.WithDetails(request)
	if err != nil {
		return nil, errors.Errorf("(*status.Status).WithDetails: %w", err)
	}

	return nil, d.Err()
}
