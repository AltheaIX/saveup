package retention

import (
	"context"
	"saveup/internal/config"
	"testing"
)

func TestRetention(t *testing.T) {
	cfg, err := config.Load("../../config.yaml")
	if err != nil {
		t.Fatalf("failed to load config: %v", err)
	}

	err = Run(context.Background(), cfg)
	if err != nil {
		t.Fatalf("failed to run: %v", err)
	}
}
