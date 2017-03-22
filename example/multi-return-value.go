package main

import "fmt"

func main() {
	a, b := vals()
	fmt.Println(a, b)

	_,c := vals()
	fmt.Println(c)
}

func vals() (int, int) {
	return 4, 6
}
