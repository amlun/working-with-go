package main

import (
	"os/exec"
	"fmt"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	dateCmd := exec.Command("date")

	dateOut, err := dateCmd.Output()
	checkErr(err)
	fmt.Println(">date")
	fmt.Println(string(dateOut))
}
