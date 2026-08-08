package storage

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

func (r *S3ClientImpl) Upload(ctx context.Context, archivePath string) error {
	file, err := os.Open(archivePath)
	if err != nil {
		return fmt.Errorf("failed to open archive: %w", err)
	}
	defer file.Close()

	key := filepath.Base(archivePath)

	fmt.Println("Uploading...")
	fmt.Println("Bucket :", r.cfg.S3.Bucket)
	fmt.Println("Object :", key)

	_, err = r.client.PutObject(
		ctx, &s3.PutObjectInput{
			Bucket: aws.String(r.cfg.S3.Bucket),
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

func (r *S3ClientImpl) List(ctx context.Context) (*s3.ListObjectsV2Output, error) {
	input := &s3.ListObjectsV2Input{
		Bucket: aws.String(r.cfg.S3.Bucket),
	}

	return r.client.ListObjectsV2(ctx, input)
}

func (r *S3ClientImpl) BulkDelete(ctx context.Context, objects []types.ObjectIdentifier) error {
	input := &s3.DeleteObjectsInput{
		Bucket: aws.String(r.cfg.S3.Bucket),
		Delete: &types.Delete{
			Objects: objects,
			Quiet:   aws.Bool(true),
		},
	}

	_, err := r.client.DeleteObjects(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to delete objects: %w", err)
	}

	return nil
}

func (r *S3ClientImpl) Check(ctx context.Context) error {
	_, err := r.client.HeadBucket(
		ctx,
		&s3.HeadBucketInput{
			Bucket: aws.String(r.cfg.S3.Bucket),
		},
	)

	if err != nil {
		return err
	}

	return nil
}
