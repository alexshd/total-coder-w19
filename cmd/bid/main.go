package main

import (
	"fmt"
	"log/slog"
	"runtime"
	"sync"
	"time"

	"github.com/alexshd/total-coder-w19/logger"
)

func main() {
	// pooling()
	boundedWorkPooling()
}

// pooling: In this pattern, the parent goroutine signals 100 pieces of work
// to a pool of child goroutines waiting for work to perform.
func pooling() {
	slog := slog.New(logger.NewLogHandler("POOLING"))
	ch := make(chan string)

	g := runtime.GOMAXPROCS(0)
	for c := range g {
		go func(child int) {
			for d := range ch {
				slog.Info("goroutine", "child", child, "recv'd signal", d)
			}
			slog.Info("shotdown singnal recived", "child", child)
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

// boundedWorkPooling: In this pattern, a pool of child goroutines is created
// to service a fixed amount of work. The parent goroutine iterates over all
// work, signaling that into the pool. Once all the work has been signaled,
// then the channel is closed, the channel is flushed, and the child
// goroutines terminate.
func boundedWorkPooling() {
	slog := slog.New(logger.NewLogHandler("POOLING"))
	work := []string{"paper1", "paper2", "paper3", "paper4"}
	slog.Info("work length", "len", len(work))
	g := runtime.GOMAXPROCS(0)
	var wg sync.WaitGroup
	wg.Add(g)

	ch := make(chan string, g)

	for c := range g {
		go func(child int) {
			defer wg.Done()
			for wrk := range ch {
				fmt.Printf("child %d : recv'd signal : %s\n", child, wrk)
			}
			fmt.Printf("child %d : recv'd shutdown signal\n", child)
		}(c)
	}

	for _, wrk := range work {
		ch <- wrk
	}
	close(ch)
	wg.Wait()

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------")
}
