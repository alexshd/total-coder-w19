package logger

import (
	"os"
	"time"

	"github.com/charmbracelet/log"
)

func NewLogHandler(prefix string) *log.Logger {
	return log.NewWithOptions(os.Stderr, log.Options{
		TimeFormat:      time.Kitchen,
		Level:           0,
		Prefix:          prefix,
		ReportTimestamp: true,
		ReportCaller:    true,
	})
}
