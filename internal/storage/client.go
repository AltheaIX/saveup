package storage

import (
	"context"
	"fmt"
	"saveup/internal/config"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type S3ClientImpl struct {
	cfg    *config.Config
	client *s3.Client
}

type S3Client interface {
	Upload(ctx context.Context, archivePath string) error
	List(ctx context.Context) (*s3.ListObjectsV2Output, error)
	BulkDelete(ctx context.Context, objects []types.ObjectIdentifier) error
	Check(ctx context.Context) error
}

func NewS3Client(cfg *config.Config) (S3Client, error) {
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

	return &S3ClientImpl{
		cfg:    cfg,
		client: client,
	}, nil
}
