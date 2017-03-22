package main

import "fmt"

func main() {
	m := make(map[string]int)
	fmt.Println(m)

	m["a"] = 1
	m["b"] = 2
	fmt.Println("map:", m)

	v1 := m["a"]
	fmt.Println("v1 is", v1)

	fmt.Println("len is", len(m))

	delete(m, "b")
	fmt.Println("map:", m)

	_, prs := m["b"]
	fmt.Println("prs", prs)

	n := map[string]int{"c": 3, "d": 4}
	fmt.Println("map:", n)
}
