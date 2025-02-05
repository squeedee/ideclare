package main

import (
	"context"
	"github.com/squeedee/ideclare/cmd"
	"os"
	"os/signal"

	"github.com/google/go-containerregistry/pkg/logs"
)

func init() {
	logs.Warn.SetOutput(os.Stderr)
	logs.Progress.SetOutput(os.Stderr)
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	if err := cmd.Root.ExecuteContext(ctx); err != nil {
		stop()
		os.Exit(1)
	}
}
