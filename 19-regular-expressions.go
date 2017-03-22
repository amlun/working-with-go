/**
 * regular-expressions.go
 *
 * A set of examples using regular expressions in Go.
 * See package documentation: http://golang.org/pkg/regexp/
 */
package main

import (
	"fmt"
	"regexp"
)

func basic_regexes() {
	// create regular expression pattern
	// pattern, match 1 or more numbers
	//pattern := "[0-9]+"
	pattern := "https://"
	// test string
	//str := "The 12 monkeys ate 48 bananas"
	str := "https://nr-cmt.ksmobile.net"
	// compile pattern
	re, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Error compiling regex", err)
	}

	// 1. Test compiled pattern matches string
	if re.MatchString(str) {
		fmt.Println("Yes, matched url scheme")
	} else {
		fmt.Println("No, no match")
	}

	// 2. Return first match
	result_two := re.FindString(str)
	fmt.Println("https matched:", result_two)

	// 3. Return n matches, use -1 to find all matches
	results_three := re.FindAllString(str, 2)
	for i, v := range results_three {
		fmt.Printf("Match %d: %s\n", i, v)
	}

	// 4. Replace matches
	results_four := re.ReplaceAllString(str, "http://")
	fmt.Println("Result:", results_four)
}

func case_insensitive() {
	// the format for flags is a bit different than typical regex
	// in golang the flag precedes the pattern, the syntax is not great
	// here is a case insensitive flag
	ptn := `(?i)^t.`
	str := "To be or not to be"

	re, err := regexp.Compile(ptn)
	if err != nil {
		fmt.Println("Error compiling regex", err)
	}

	// match string
	result := re.FindString(str)
	fmt.Println("Result:", result)

	results := re.FindAllString(str, 4)
	for i, v := range results {
		fmt.Printf("Match %d: %s\n", i, v)
	}
}

func main() {
	basic_regexes()
	case_insensitive()
}
