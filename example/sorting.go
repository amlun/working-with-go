package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"c", "a", "b"}
	fmt.Println(strs)
	sort.Strings(strs)
	fmt.Println(strs)

	ints := []int{1, 4, 2, 7, 5}
	fmt.Println(ints)
	fmt.Println(sort.IntsAreSorted(ints))

	sort.Ints(ints)
	fmt.Println(ints)
	fmt.Println(sort.IntsAreSorted(ints))

}
