package main

import "fmt"

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(addNew(1, 2, 3))
}

func add(a int, b int) int {
	return a + b
}

func addNew(a, b, c int) int {
	return a + b + c
}
