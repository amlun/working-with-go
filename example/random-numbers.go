package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	p := fmt.Print
	pl := fmt.Println

	p(rand.Intn(100), ",")
	p(rand.Intn(100))
	pl()

	pl(rand.Float64())

	p((rand.Float64()*5)+5, ",")
	p((rand.Float64() * 5) + 5)
	pl()


}
