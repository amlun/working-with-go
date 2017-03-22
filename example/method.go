package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.height + 2*r.width
}

func main() {
	r := rect{3, 4}
	fmt.Println(r)
	fmt.Println("perim: ", r.perim())
	fmt.Println("area: ", r.area())

	rp := &r
	fmt.Println("perim: ", rp.perim())
	fmt.Println("area: ", rp.area())

}
