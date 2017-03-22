package main

import (
	"time"
	"fmt"
)

func main() {
	c1 := make(chan string, 1)

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "result 1"
	}()

	select {
	case r1 := <-c1:
		fmt.Println("received from r1:", r1)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 1")

	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "result 2"
	}()

	select {
	case r2 := <-c2:
		fmt.Println("received from r2:", r2)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout 2")
	}
}
