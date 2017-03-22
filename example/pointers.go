package main

import "fmt"

func main() {
	ival := 1
	fmt.Println("ival=", ival)
	zeroval(ival)
	fmt.Println("ival=", ival)

	zeroptr(&ival)
	fmt.Println("ival=", ival)

	fmt.Println("&ival=", &ival)

}

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}
