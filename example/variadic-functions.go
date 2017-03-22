package main

import "fmt"

func main() {
	fmt.Println(sum(1, 2, 3, 3, 4, 5, 6, 7))
	// If you already have multiple args in a slice, apply them to a variadic function using func(slice...) like this.
	nums := []int{1, 2, 3, 4}
	fmt.Println(sum(nums...))
}

func sum(nums ...int) int {
	fmt.Println(nums)
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
