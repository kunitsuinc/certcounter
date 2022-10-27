package controller_test

import (
	"context"
	"testing"

	v1 "github.com/kunitsuinc/certcounter/generated/go/certcounter/v1"
	"github.com/kunitsuinc/certcounter/pkg/controller"
)

func TestTestAPIController_Echo(t *testing.T) {
	t.Parallel()
	t.Run("success()", func(t *testing.T) {
		t.Parallel()
		tr := &controller.TestAPIController{}
		const wantMessage = "test"
		gotResponse, err := tr.Echo(context.Background(), &v1.TestAPIServiceEchoRequestResponse{Message: wantMessage})
		if err != nil {
			t.Errorf("TestAPIController.Echo() error = %v, wantErr %v", err, nil)
			return
		}
		if gotResponse.Message != wantMessage {
			t.Errorf("TestAPIController.Echo(): Message = %v, want %v", gotResponse.Message, wantMessage)
		}
	})
}
