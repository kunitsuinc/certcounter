package controller

import (
	"context"

	"github.com/kunitsuinc/certcounter/pkg/traces"
	testapiv1 "github.com/kunitsuinc/certcounter/proto/generated/go/testapi/v1"
)

type TestAPIController struct {
	testapiv1.UnimplementedTestAPIServiceServer
}

func (*TestAPIController) Echo(ctx context.Context, request *testapiv1.TestAPIServiceEchoRequest) (response *testapiv1.TestAPIServiceEchoResponse, err error) {
	_ = traces.StartFunc(ctx, "Echo")(func(ctx context.Context) error {
		response = &testapiv1.TestAPIServiceEchoResponse{
			Message: request.GetMessage(),
		}

		return nil
	})

	return response, nil
}
