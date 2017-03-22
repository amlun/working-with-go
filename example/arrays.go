package main

import (
	"fmt"
)

func main() {
	fmt.Println("OK")
	// a is a int array with 5 items
	var a [5]int

	fmt.Println(a)

	a[4] = 100
	fmt.Println(a)

	fmt.Println(len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	var twoD [5][5]int
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)
}
