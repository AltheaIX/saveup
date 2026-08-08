package diagnostic

import (
	"context"
	"fmt"
	"os"
	"saveup/internal/storage"

	"saveup/internal/config"
)

type Diagnostic struct {
	Client storage.S3Client
	Cfg    *config.Config
}

func (d *Diagnostic) Run(ctx context.Context, cfg *config.Config) error {
	_, err := os.Stat(cfg.Save.Path)
	if err != nil {
		return fmt.Errorf("save directory not found: %s", cfg.Save.Path)
	}

	fmt.Println("✓ Validate directory exists")

	fmt.Println()

	if err = d.Client.Check(ctx); err != nil {
		return err
	}
	fmt.Println("✓ Validate S3 connection")
	fmt.Println()

	fmt.Println("PASS")

	return nil
}
