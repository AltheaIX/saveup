package backup

import (
	"fmt"
	"os"
	"path/filepath"
	"saveup/internal/compression"
	"saveup/internal/config"
	"slices"
	"strings"
	"time"
)

func LatestArchive(workspace string) (string, error) {
	entries, err := os.ReadDir(workspace)
	if err != nil {
		return "", err
	}

	var archives []string

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		if strings.HasSuffix(entry.Name(), ".zip") {
			archives = append(archives, entry.Name())
		}
	}

	if len(archives) == 0 {
		return "", fmt.Errorf("no archive found")
	}

	slices.Sort(archives)

	return filepath.Join(workspace, archives[len(archives)-1]), nil
}

func generateFileName(prefix, ext string) string {
	return fmt.Sprintf(
		"%s-%s.%s",
		prefix,
		time.Now().Format("20060102-150405"),
		ext,
	)
}

func Run(cfg *config.Config) (string, error) {
	if err := os.MkdirAll(cfg.Workspace.Path, 0755); err != nil {
		return "", fmt.Errorf("failed to create workspace: %w", err)
	}

	archivePath := filepath.Join(
		cfg.Workspace.Path,
		generateFileName(cfg.Backup.Prefix, "zip"),
	)

	sourcePath := filepath.Join(cfg.Save.Path)

	fmt.Println("Creating backup...")
	fmt.Println("Source: ", sourcePath)
	fmt.Println("Output: ", archivePath)
	fmt.Println()

	err := compression.Zip(sourcePath, archivePath)
	if err != nil {
		return "", fmt.Errorf("failed to compress save directory: %w", err)
	}

	fmt.Println("✓ Backup completed")

	return archivePath, nil
}
