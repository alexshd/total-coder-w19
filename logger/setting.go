package logger

import (
	"os"
	"time"

	"github.com/charmbracelet/log"
)

var Log = log.NewWithOptions(os.Stderr, log.Options{
	TimeFormat:      time.RFC3339Nano,
	Level:           0,
	Prefix:          "",
	ReportTimestamp: true,
	ReportCaller:    true,
	CallerOffset:    2,
})
