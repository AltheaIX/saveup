package daemon

import (
	"context"
	"fmt"
	"saveup/internal/backup"
	"saveup/internal/config"
	"saveup/internal/storage"
	"time"
)

func Run(ctx context.Context, cfg *config.Config) error {
	duration, err := time.ParseDuration(cfg.Backup.Interval)
	if err != nil {
		return fmt.Errorf("parse duration: %w", err)
	}

	ticker := time.NewTicker(duration)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			archivePath, err := backup.Run(cfg)
			if err != nil {
				return fmt.Errorf("backup: %w", err)
			}

			if err = runUpload(cfg, archivePath); err != nil {
				return fmt.Errorf("upload: %w", err)
			}

			if !cfg.Backup.KeepLocal {
				err = backup.DeleteFile(archivePath)
				if err != nil {
					return fmt.Errorf("deleting uploaded file: %w", err)
				}
			}
		case <-ctx.Done():
			fmt.Println("Received shutdown signal")
			return nil
		}
	}
}

func runUpload(cfg *config.Config, archivePath string) error {
	// using uploadCtx to prevent backup getting canceled mid-way
	uploadCtx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	return storage.Upload(uploadCtx, cfg, archivePath)
}
