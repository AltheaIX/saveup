package diagnostic

import (
	"context"
	"fmt"
	"os"
	"saveup/internal/storage"

	"saveup/internal/config"
)

func Run(ctx context.Context, cfg *config.Config) error {
	_, err := os.Stat(cfg.Save.Path)
	if err != nil {
		return fmt.Errorf("save directory not found: %s", cfg.Save.Path)
	}

	fmt.Println("✓ Validate directory exists")

	fmt.Println()

	if err = storage.Check(ctx, cfg); err != nil {
		return err
	}
	fmt.Println("✓ Validate S3 connection")
	fmt.Println()

	fmt.Println("PASS")

	return nil
}
