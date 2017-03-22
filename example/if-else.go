package main

import "fmt"

func main() {

	n := 17

	if isEven(n) {
		fmt.Printf("%d is even\n", n)
	} else {
		fmt.Printf("%d is odd\n", n)
	}

	if m := sub(n); m > 0 {
		fmt.Printf("%d is more than 10\n", n)
	}

}

func sub(n int) int {
	return n - 10
}

func isEven(n int) bool {
	if n%2 == 0 {
		return true
	} else {
		return false
	}
}
