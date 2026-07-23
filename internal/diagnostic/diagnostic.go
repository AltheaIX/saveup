package diagnostic

import (
	"fmt"
	"os"

	"saveup/internal/config"
)

func Run() error {

	cfg, err := config.Load("config.yaml")
	if err != nil {
		return err
	}

	fmt.Println("✓ config.yaml parsed")

	_, err = os.Stat(cfg.Save.Path)
	if err != nil {
		return fmt.Errorf("save directory not found: %s", cfg.Save.Path)
	}

	fmt.Println("✓ save directory exists")

	fmt.Println()
	fmt.Println("PASS")

	return nil
}
