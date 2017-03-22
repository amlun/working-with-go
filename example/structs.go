package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"lunweiwei", 28})

	fmt.Println(person{"Alice", 30})

	fmt.Println(person{name: "Job", age: 50})

	fmt.Println(&person{"Bosh", 55})

	s := person{"Sean", 19}
	fmt.Println(s)

	sp := &s
	fmt.Println(sp.name)

	sp.name = "Sean1"
	fmt.Println(s)
	fmt.Println(sp)
}
