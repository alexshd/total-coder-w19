package main

import (
	"fmt"
	"log/slog"
	"runtime"
	"sync"
	"time"

	"github.com/alexshd/total-coder-w19/internal/logger"
)

func main() {
	Pooling()
	BoundedWorkPooling()
}

// Pooling: In this pattern, the parent goroutine signals 100 pieces of work
// to a pool of child goroutines waiting for work to perform.
func Pooling() {
	slog := slog.New(logger.NewLogHandler("POOLING-A"))
	ch := make(chan string)

	g := runtime.GOMAXPROCS(0)
	for c := range g {
		go func(child int) {
			for d := range ch {
				slog.Info("goroutine", "child", child, "recv'd signal", d)
			}
			slog.Info("shutdown signal received", "child", child)
		}(c)
	}

	const work = 100
	for w := 0; w < work; w++ {
		ch <- "data"
		slog.Info("parent", "sent signal", w)
	}

	close(ch)
	slog.Info("parent : sent shutdown signal")

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------")
}

// BoundedWorkPooling: In this pattern, a pool of child goroutines is created
// to service a fixed amount of work. The parent goroutine iterates over all
// work, signaling that into the pool. Once all the work has been signaled,
// then the channel is closed, the channel is flushed, and the child
// goroutines terminate.
func BoundedWorkPooling() {
	var (
		slog = slog.New(logger.NewLogHandler("POOLING-B"))

		work = []string{"paper1", "paper2", "paper3", "paper4"}
		g    = runtime.GOMAXPROCS(0)
		wg   sync.WaitGroup
	)

	wg.Add(g)

	ch := make(chan string, g)

	for c := range g {
		go func(child int) {
			defer wg.Done()
			for wr := range ch {
				slog.Info("work +", "child", child, "recv'd signal", wr)
			}
			slog.Info("work -", "child", child)
		}(c)
	}

	for _, wr := range work {
		ch <- wr
	}
	close(ch)
	wg.Wait()

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------")
}
