package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func main() {
	done := make(chan bool)

	// start a worker goroutine
	go worker(done)

	// block until receive a notification from the worker on the channel
	<-done
}
