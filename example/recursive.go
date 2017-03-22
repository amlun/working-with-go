package main

import "fmt"

func Top(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}
	return Top(n-1) + Top(n-2)
}

func main() {
	fmt.Println(1, Top(1))
	fmt.Println(2, Top(2))
	fmt.Println(3, Top(3))
	fmt.Println(4, Top(4))
	fmt.Println(5, Top(5))
}
