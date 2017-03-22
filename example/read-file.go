package main

import (
	"io/ioutil"
	"fmt"
	"os"
	"strings"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	data, err := ioutil.ReadFile("/tmp/data")
	checkError(err)
	fmt.Println(string(data))

	f, err := os.Open("/tmp/data")
	defer f.Close()
	checkError(err)

	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	checkError(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))

	d1 := []byte("\ngo\nhello\n")
	err = ioutil.WriteFile("/tmp/data", d1, 0644)
	checkError(err)

	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0])
	}
}
