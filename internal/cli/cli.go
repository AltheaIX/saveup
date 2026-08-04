package cli

import (
	"context"
	"fmt"
	"os"
	"saveup/internal/backup"
	"saveup/internal/config"
	"saveup/internal/daemon"
	"saveup/internal/diagnostic"
	"saveup/internal/storage"
)

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
	fmt.Println()

	switch os.Args[1] {
	case "diag":
		if err = diagnostic.Run(ctx, cfg); err != nil {
			return fmt.Errorf("run diagnostic: %w\n", err)
		}
	case "backup":
		if _, err = backup.Run(cfg); err != nil {
			return fmt.Errorf("run backup: %w\n", err)
		}
	case "store":
		return runStore(ctx, cfg)
	case "daemon":
		if err = diagnostic.Run(ctx, cfg); err != nil {
			return fmt.Errorf("run diagnostic: %w\n", err)
		}

		err = daemon.Run(ctx, cfg)
		if err != nil {
			return fmt.Errorf("running daemon: %w\n", err)
		}
	default:
		help()
	}

	return nil
}

func runStore(ctx context.Context, cfg *config.Config) (err error) {
	if len(os.Args) < 3 {
		fmt.Println("Usage: saveup store <archive>")
		return nil
	}

	archiveFile := os.Args[2]
	if archiveFile == "latest" {
		archiveFile, err = backup.LatestArchive(cfg.Workspace.Path)
		if err != nil {
			return fmt.Errorf("getting latest archive: %w\n", err)
		}
	}

	if err = storage.Upload(ctx, cfg, archiveFile); err != nil {
		return fmt.Errorf("running upload: %w\n", err)
	}

	if !cfg.Backup.KeepLocal {
		return backup.DeleteFile(archiveFile)
	}

	return nil
}

func help() {
	fmt.Println("saveup")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  saveup diag")
	fmt.Println("  saveup backup")
	fmt.Println("  saveup store [file] - use absolute path from backup's output")
	fmt.Println("  # Use `saveup store latest` to store the latest version after running `saveup backup`")
	fmt.Println("  saveup daemon - start daemon for auto-backup")
}
