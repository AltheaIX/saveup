package storage

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"saveup/internal/config"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func Upload(ctx context.Context, cfg *config.Config, archivePath string) error {
	file, err := os.Open(archivePath)
	if err != nil {
		return fmt.Errorf("failed to open archive: %w", err)
	}
	defer file.Close()

	key := filepath.Base(archivePath)

	fmt.Println("Uploading...")
	fmt.Println("Bucket :", cfg.S3.Bucket)
	fmt.Println("Object :", key)

	client, err := NewR2Client(cfg)
	if err != nil {
		return fmt.Errorf("failed to create R2 client: %w", err)
	}

	_, err = client.PutObject(
		ctx, &s3.PutObjectInput{
			Bucket: aws.String(cfg.S3.Bucket),
			Key:    aws.String(key),
			Body:   file,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to upload object: %w", err)
	}

	fmt.Println("✓ Upload completed")
	return nil
}

func Check(ctx context.Context, cfg *config.Config) error {
	client, err := NewR2Client(cfg)
	if err != nil {
		return err
	}

	_, err = client.HeadBucket(
		ctx,
		&s3.HeadBucketInput{
			Bucket: aws.String(cfg.S3.Bucket),
		},
	)

	if err != nil {
		return err
	}

	return nil
}
