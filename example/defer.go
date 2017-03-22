package main

import (
	"os"
	"fmt"
)

func main() {
	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(filename string) *os.File {
	fmt.Println("creating...", filename)

	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing...")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing...")
	f.Close()
}
