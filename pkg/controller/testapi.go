package controller

import (
	"context"

	statusz "github.com/kunitsuinc/grpcutil.go/grpc/status"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/emptypb"

	v1 "github.com/kunitsuinc/certcounter/generated/go/certcounter/v1"
	"github.com/kunitsuinc/certcounter/pkg/errors"
	"github.com/kunitsuinc/certcounter/pkg/traces"
)

type TestAPIController struct {
	v1.UnimplementedTestAPIServiceServer
}

func (*TestAPIController) Echo(ctx context.Context, request *v1.TestAPIServiceEchoRequestResponse) (*v1.TestAPIServiceEchoRequestResponse, error) {
	_, span := traces.Start(ctx, "Echo")
	defer span.End()

	if err := request.Validate(); err != nil {
		return nil, statusz.New(codes.InvalidArgument, "InvalidArgument: "+err.Error(), errors.Errorf("request.Validate: %w", err))
	}

	return request, nil
}

func (*TestAPIController) EchoError(ctx context.Context, request *v1.TestAPIServiceEchoErrorRequest) (*emptypb.Empty, error) {
	_, span := traces.Start(ctx, "EchoError")
	defer span.End()

	if err := request.Validate(); err != nil {
		return nil, statusz.New(codes.InvalidArgument, "InvalidArgument: "+err.Error(), errors.Errorf("request.Validate: %w", err))
	}

	s := statusz.New(codes.Code(request.GetCode()), request.GetMessage(), errors.Errorf(request.GetMessage()))

	d, err := s.WithDetails(request)
	if err != nil {
		return nil, statusz.New(codes.Internal, "Internal: "+err.Error(), errors.Errorf("(*status.Status).WithDetails: %w", err))
	}

	return nil, d
}
