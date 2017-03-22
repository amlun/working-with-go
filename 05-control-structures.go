/**
 * control-structures.go
 * If, Else, Switch and Conditionals
 */

// standard stuff
package main

import "fmt"

func main() {

	// define an int variable
	num := 1

	// Go is not picky, conditionals don't require parentheses
	if num > 3 {
		fmt.Println("Many")
	}

	// Go is picky, "else" must be on the same line as closing if bracket
	if num == 1 {
		fmt.Println("One")
	} else if num == 2 {
		fmt.Println("Two")
	} else {
		fmt.Println("Many")
	}

	// Switch statement takes conditionals and auto breaks
	switch {
	case num == 1:
		fmt.Println("One")
	case num == 2:
		fmt.Println("Two")
	case num > 2:
		fmt.Println("Many")
	default:
		fmt.Println("Thrown over boat")
	}

	weekDay(-1)
	weekDay(1)
	weekDay(6)
	weekDay(18)

}

func weekDay(day int) {
	switch {
	case day == 1:
		fmt.Println("星期一")
	case day == 2:
		fmt.Println("星期二")
	case day == 3:
		fmt.Println("星期三")
	case day == 4:
		fmt.Println("星期四")
	case day == 5:
		fmt.Println("星期五")
	case day > 7 || day < 1:
		fmt.Println("Error")
	default:
		fmt.Println("周末")
	}
}
