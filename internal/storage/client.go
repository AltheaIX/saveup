package storage

import (
	"context"
	"fmt"
	"saveup/internal/config"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func NewR2Client(cfg *config.Config) (*s3.Client, error) {
	awscfg, err := awsconfig.LoadDefaultConfig(
		context.TODO(),
		awsconfig.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(
				cfg.S3.AccessKey,
				cfg.S3.SecretKey,
				"",
			),
		),
		awsconfig.WithRegion("auto"), // Required by SDK but not used by R2
	)

	if err != nil {
		return nil, fmt.Errorf("failed loading aws config: %w", err)
	}

	client := s3.NewFromConfig(
		awscfg, func(o *s3.Options) {
			o.BaseEndpoint = aws.String(cfg.S3.Endpoint)
		},
	)

	return client, nil
}
