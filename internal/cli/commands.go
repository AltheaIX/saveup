package cli

import (
	"context"
	"fmt"
	"os"
	"saveup/internal/backup"
	"saveup/internal/config"
	"saveup/internal/daemon"
	"saveup/internal/diagnostic"
	"saveup/internal/list"
	"saveup/internal/retention"
	"saveup/internal/storage"
)

func runDiagnostic(ctx context.Context, cfg *config.Config, client storage.S3Client) error {
	diag := &diagnostic.Diagnostic{
		Client: client,
		Cfg:    cfg,
	}
	return diag.Run(ctx, cfg)
}

func runList(ctx context.Context, cfg *config.Config, client storage.S3Client) error {
	listImpl := &list.List{
		Client: client,
	}

	return listImpl.Run(ctx)
}

func runBackup(ctx context.Context, cfg *config.Config, client storage.S3Client) error {
	if _, err := backup.Run(cfg); err != nil {
		return fmt.Errorf("backup failed: %w", err)
	}

	return nil
}

func runStore(ctx context.Context, cfg *config.Config, client storage.S3Client) (err error) {
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

	if err = client.Upload(ctx, archiveFile); err != nil {
		return fmt.Errorf("running upload: %w\n", err)
	}

	if !cfg.Backup.KeepLocal {
		return backup.DeleteFile(archiveFile)
	}

	return nil
}

func runDaemon(ctx context.Context, cfg *config.Config, client storage.S3Client) error {
	ret := &retention.Retention{
		Client: client,
	}

	dae := &daemon.Daemon{
		Retention: ret,
		Cfg:       cfg,
		Client:    client,
	}

	diag := &diagnostic.Diagnostic{
		Client: client,
		Cfg:    cfg,
	}

	if err := diag.Run(ctx, cfg); err != nil {
		return fmt.Errorf("running diagnostic: %w", err)
	}

	if err := dae.Run(ctx); err != nil {
		return fmt.Errorf("running daemon: %w", err)
	}

	return nil
}
