package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"saveup/internal/cli"
	"syscall"
)

func main() {
	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	if err := cli.Execute(ctx); err != nil {
		log.Fatal(err)
	}
}
