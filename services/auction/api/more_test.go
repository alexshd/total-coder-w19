package api

import (
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alexshd/total-coder-w19/logger"
	"github.com/stretchr/testify/assert"
)

func TestSomeTest(t *testing.T) {
	a := assert.New(t)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, client")
	}))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}
	greeting, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	a.Equal("Hello, client\n", string(greeting))
	fmt.Printf("%s", greeting)
}

func TestGETPlayers(t *testing.T) {
	t.Run("returns Pepper's score", func(t *testing.T) {
		a := assert.New(t)
		request, _ := http.NewRequest(http.MethodGet, "/players/Pepper", nil)

		response := httptest.NewRecorder()

		PlayerServer(response, request)

		got := response.Body.String()
		want := "20"

		a.Equal(want, got)
	})
}

func HttpLogResponse(r *http.Request) {
	slog := slog.New(logger.NewLogHandler("HTTP"))
	for key, val := range r.Header {
		slog.Info("Header", key, val)
	}
}
