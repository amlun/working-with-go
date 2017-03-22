package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)
	go func() { messages <- "ping" }()

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("send message", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received from messages: ", msg)
	default:
		fmt.Println("no message received")
	}

	go func() { messages <- "pong" }()
	select {
	case msg := <-messages:
		fmt.Println("received messages:", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
