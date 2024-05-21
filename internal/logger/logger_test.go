package logger

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/charmbracelet/log"
)

func TestLogger(t *testing.T) {
	t.Parallel()

	t.Run("NewLogHandler is of type log.Logger", func(t *testing.T) {
		got := reflect.TypeOf(NewLogHandler("Test"))
		want := reflect.TypeOf(new(log.Logger))

		if got != want {
			t.Fatalf("Expected: %T , Got: %T", want, got)
		}
	})
}

func TestA(t *testing.T) {
	t.Log("Test A run")
}

func TestB(t *testing.T) {
	t.Log("Test B run")
}

func TestMain(m *testing.M) {
	exitVal := m.Run()

	if exitVal == 0 {
		// teardown()
		fmt.Println("Teardown")
	}
	os.Exit(exitVal)
}
