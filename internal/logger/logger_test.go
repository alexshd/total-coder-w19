package logger

import (
	"testing"

	"github.com/charmbracelet/log"
	"github.com/stretchr/testify/assert"
)

func TestLogger(t *testing.T) {
	a := assert.New(t)

	handler := NewLogHandler("Test")

	a.IsType(new(log.Logger), handler)
}
