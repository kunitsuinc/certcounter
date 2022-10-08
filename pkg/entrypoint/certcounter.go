package entrypoint

import (
	"context"

	"github.com/kunitsuinc/certcounter/pkg/config"
)

func CertCounter(ctx context.Context) error {
	gcpProjectID := config.GoogleCloudProject()
	awsProfile := config.AWSProfile()

	_ = gcpProjectID
	_ = awsProfile

	return nil
}
