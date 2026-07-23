package backup

import (
	"fmt"
	"os"
	"path/filepath"
	"saveup/internal/compression"
	"saveup/internal/config"
	"time"
)

func generateFileName(prefix, ext string) string {
	return fmt.Sprintf(
		"%s-%s.%s",
		prefix,
		time.Now().Format("20060102-150405"),
		ext,
	)
}

func Run() (string, error) {
	fmt.Println("Loading configuration...")
	cfg, err := config.Load("config.yaml")
	if err != nil {
		return "", err
	}
	fmt.Println("✓ Configuration loaded")
	fmt.Println()

	if err = os.MkdirAll(cfg.Workspace.Path, 0755); err != nil {
		return "", fmt.Errorf("failed to create workspace: %w", err)
	}

	archivePath := filepath.Join(
		cfg.Workspace.Path,
		generateFileName("palworld", "zip"),
	)

	sourcePath := filepath.Join(cfg.Save.Path)

	fmt.Println("Creating backup...")
	fmt.Println("Source: ", sourcePath)
	fmt.Println("Output: ", archivePath)
	fmt.Println()

	err = compression.Zip(sourcePath, archivePath)
	if err != nil {
		return "", fmt.Errorf("failed to compress save directory: %w", err)
	}

	fmt.Println("✓ Backup completed")

	return archivePath, nil
}
