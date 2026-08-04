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

	err = Upload(context.Background(), cfg, "D:\\Palworld\\Workspace\\palworld-20260723-222347.zip")
	if err != nil {
		log.Fatal(err)
	}
}
