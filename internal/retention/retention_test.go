package retention

import (
	"context"
	"saveup/internal/config"
	"saveup/internal/storage"
	"testing"
)

func TestRetention(t *testing.T) {
	cfg, err := config.Load("../../config.yaml")
	if err != nil {
		t.Fatalf("failed to load config: %v", err)
	}

	client, err := storage.NewS3Client(cfg)
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}

	retention := &Retention{
		Client: client,
	}

	err = retention.Run(context.Background())
	if err != nil {
		t.Fatalf("failed to run: %v", err)
	}
}
