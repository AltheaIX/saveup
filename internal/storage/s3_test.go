package storage

import (
	"context"
	"log"
	"saveup/internal/config"
	"testing"
)

func TestUpload(t *testing.T) {
	cfg, err := config.Load("../../config.yaml")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	client, err := NewS3Client(cfg)
	if err != nil {
		log.Fatalf("failed to create S3 client: %v", err)
	}

	err = client.Upload(context.Background(), "D:\\Palworld\\Workspace\\palworld-20260723-222347.zip")
	if err != nil {
		log.Fatal(err)
	}
}
