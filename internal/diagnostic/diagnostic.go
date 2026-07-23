package diagnostic

import (
	"fmt"
	"os"
	"saveup/internal/storage"

	"saveup/internal/config"
)

func Run(cfg *config.Config) error {
	_, err := os.Stat(cfg.Save.Path)
	if err != nil {
		return fmt.Errorf("save directory not found: %s", cfg.Save.Path)
	}

	fmt.Println("✓ save directory exists")

	fmt.Println()

	if err = storage.Check(cfg); err != nil {
		return err
	}
	fmt.Println("✓ R2 connection")
	fmt.Println()

	fmt.Println("PASS")

	return nil
}
