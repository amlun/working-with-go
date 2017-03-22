package main

import (
	"crypto/sha1"
	"fmt"
	"crypto/md5"
)

func main() {
	s := "sha1 this string"

	h := sha1.New()
	h.Write([]byte(s))

	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)

	h1 := md5.New()
	h1.Write([]byte(s))

	bs1 := h1.Sum(nil)
	fmt.Printf("%x\n", bs1)
}
