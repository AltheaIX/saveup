package daemon

import (
	"context"
	"fmt"
	"saveup/internal/backup"
	"saveup/internal/config"
	"saveup/internal/retention"
	"saveup/internal/storage"
	"time"
)

type Daemon struct {
	Retention *retention.Retention

	Cfg    *config.Config
	Client storage.S3Client
}

func (d *Daemon) Run(ctx context.Context) error {
	duration, err := time.ParseDuration(d.Cfg.Backup.Interval)
	if err != nil {
		return fmt.Errorf("parse duration: %w", err)
	}

	ticker := time.NewTicker(duration)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			err = d.runAutoBackup()
			if err != nil {
				return err
			}

			err = d.Retention.Run(ctx)
			if err != nil {
				return err
			}
		case <-ctx.Done():
			fmt.Println("Received shutdown signal")
			return nil
		}
	}
}

func (d *Daemon) runAutoBackup() error {
	archivePath, err := backup.Run(d.Cfg)
	if err != nil {
		return fmt.Errorf("backup: %w", err)
	}

	if err = d.runUpload(archivePath); err != nil {
		return fmt.Errorf("upload: %w", err)
	}

	if !d.Cfg.Backup.KeepLocal {
		err = backup.DeleteFile(archivePath)
		if err != nil {
			return fmt.Errorf("deleting uploaded file: %w", err)
		}
	}
	return nil
}

func (d *Daemon) runUpload(archivePath string) error {
	// using uploadCtx to prevent backup getting canceled mid-way
	uploadCtx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	return d.Client.Upload(uploadCtx, archivePath)
}
