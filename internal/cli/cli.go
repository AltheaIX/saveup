package cli

import (
	"fmt"
	"os"
	"saveup/internal/diagnostic"
)

func Execute() {
	if len(os.Args) < 2 {
		help()
		return
	}

	switch os.Args[1] {

	case "diag":
		if err := diagnostic.Run(); err != nil {
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
}
