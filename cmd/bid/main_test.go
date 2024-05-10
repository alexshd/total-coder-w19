package main

import (
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	log := slog.New(handler)
	a := assert.New(t)
	err := run(log)

	a.NoError(err)
}
