package main

import (
	"bufio"
	"log/slog"
	"os"
	"sync"

	"github.com/alexshd/total-coder-w19/logger"
	"github.com/gen2brain/beeep"
)

func readData(file string) <-chan string {
	log := slog.New(logger.Log)
	f, err := os.Open(file) // opens the file for reading
	if err != nil {
		log.Error("failed", "err", err)
	}

	out := make(chan string) // channel declared

	// returns a scanner to read from f
	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines) // scanning it line-by-line token

	// loop through the fileScanner based on our token split
	go func() {
		for fileScanner.Scan() {
			val := fileScanner.Text() // returns the recent token
			out <- val                // passed the token value to our channel
		}

		close(out) // closed the channel when all content of file is read

		// closed the file
		err := f.Close()
		if err != nil {
			log.Error("Unable to close an opened file", "err", err)
			return
		}
	}()

	return out
}

func fanInMergeData(ch1, ch2 <-chan string) chan string {
	chRes := make(chan string)
	var wg sync.WaitGroup
	wg.Add(2)

	// reads from 1st channel
	go func() {
		for val := range ch1 {
			chRes <- val
		}
		wg.Done()
	}()

	// reads from 2nd channel
	go func() {
		for val := range ch2 {
			chRes <- val
		}
		wg.Done()
	}()

	go func() {
		wg.Wait()    // waits till the goroutines are completed and wg marked Done
		close(chRes) // close the result channel
	}()

	return chRes
}

func main() {
	// show a notification
	beeep.Notify("DSP-TEST", "PASSED", "/home/alex/Downloads/square-check-solid.svg")
	// show a notification and play a alert sound
	beeep.Alert("DSP-TEST", "FAILED", "/home/alex/Downloads/circle-exclamation-solid.svg")
	slog := slog.New(logger.Log)
	ch1 := readData("text1.txt")
	ch2 := readData("text2.txt")

	// receive data from multiple channels and place it on result channel - FanIn
	chRes := fanInMergeData(ch1, ch2)

	// some logic with the result channel
	for val := range chRes {
		slog.Info(val)
	}
}
