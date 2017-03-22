package main

import (
	"fmt"
	"time"
)

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	// worker goroutine
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
				time.Sleep(time.Second * 2)
				fmt.Println("done job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}

	close(jobs)
	fmt.Println("send all jobs")

	<-done
	fmt.Println("done all jobs")
}
