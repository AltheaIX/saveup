package cli

import (
	"fmt"
	"os"
	"saveup/internal/backup"
	"saveup/internal/config"
	"saveup/internal/diagnostic"
	"saveup/internal/storage"
)

func Execute() {
	if len(os.Args) < 2 {
		help()
		return
	}

	fmt.Println("Loading configuration...")
	cfg, err := config.Load("config.yaml")
	if err != nil {
		fmt.Printf("Failed to load config: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("✓ Configuration loaded")
	fmt.Println()

	switch os.Args[1] {

	case "diag":
		if err = diagnostic.Run(cfg); err != nil {
			fmt.Println("ERROR:", err)
			os.Exit(1)
		}
	case "backup":
		if _, err = backup.Run(cfg); err != nil {
			fmt.Println("ERROR:", err)
			os.Exit(1)
		}
	case "store":
		if err = storage.Upload(cfg, os.Args[2]); err != nil {
			fmt.Println("ERROR:", err)
			os.Exit(1)
		}
	default:
		help()
	}
}

func help() {
	fmt.Println("saveup")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  saveup diag")
	fmt.Println("  saveup backup")
	fmt.Println("  saveup store")
}
