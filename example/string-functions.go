package main

import "fmt"
import s "strings"

var p = fmt.Println

func main() {
	p("Contains: ", s.Contains("test", "es"))
	p("Count: ", s.Count("test", "t"))
	p("HasPrefix: ", s.HasPrefix("test", "t"))
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	p("Index: ", s.Index("test", "e"))
	p("Char: ", "hello"[1])
}
