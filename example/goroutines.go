package main

import (
	"fmt"
	"math/rand"
	"time"
)

func f(from string) {
	rand.Seed(time.Now().Unix())
	num := rand.Intn(10)
	for i := 0; i < num; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	go f("goroutine1")
	go f("goroutine2")
	go f("goroutine3")
	go f("goroutine4")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
