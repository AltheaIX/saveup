package cli

import (
	"context"
	"fmt"
	"os"
	"saveup/internal/config"
	"saveup/internal/storage"
)

type commandFunc func(ctx context.Context, cfg *config.Config, client storage.S3Client) error

func Execute(ctx context.Context) error {
	if len(os.Args) < 2 {
		help()
		return nil
	}

	switch os.Args[1] {
	case "help", "-h", "--help":
		help()
		return nil
	}

	fmt.Println("Loading configuration...")
	cfg, err := config.Load("config.yaml")
	if err != nil {
		return fmt.Errorf("load config: %w\n", err)
	}
	fmt.Println("✓ Validate config")

	s3Client, err := storage.NewS3Client(cfg)
	if err != nil {
		return fmt.Errorf("create S3 client: %w\n", err)
	}
	fmt.Println()

	commands := map[string]commandFunc{
		"diag":   runDiagnostic,
		"list":   runList,
		"backup": runBackup,
		"store":  runStore,
		"daemon": runDaemon,
	}

	cmd := os.Args[1]
	if fn, ok := commands[cmd]; ok {
		return fn(ctx, cfg, s3Client)
	}

	help()
	return nil
}

func help() {
	fmt.Println("saveup")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  saveup diag")
	fmt.Println("  saveup list")
	fmt.Println("  saveup backup")
	fmt.Println("  saveup store [file] - use absolute path from backup's output")
	fmt.Println("  # Use `saveup store latest` to store the latest version after running `saveup backup`")
	fmt.Println("  saveup daemon - start daemon for auto-backup")
}
