package list

import (
	"context"
	"saveup/internal/config"
	"saveup/internal/storage"
	"testing"
)

func TestList_Run(t *testing.T) {
	cfg, err := config.Load("../../config.yaml")
	if err != nil {
		t.Fatalf("failed to load config: %v", err)
	}

	s3Client, err := storage.NewS3Client(cfg)
	if err != nil {
		t.Fatalf("failed to create S3 client: %v", err)
	}

	listImpl := &List{
		Client: s3Client,
	}

	err = listImpl.Run(context.Background())
	if err != nil {
		t.Fatalf("failed to run: %v", err)
	}
}
