package main

import (
	"log/slog"
	"os"
	"time"

	"github.com/charmbracelet/log"
)

var handler = log.NewWithOptions(os.Stderr, log.Options{
	ReportTimestamp: true,
	ReportCaller:    true,
	TimeFormat:      time.RFC3339Nano,
})

func main() {
	slog := slog.New(handler)

	if err := run(slog); err != nil {
		log.Fatal("bidding service failed", "err", err)
	}
}

func run(slog *slog.Logger) error {
	slog.Info("Hello from run", "the", "best")
	return nil
}
