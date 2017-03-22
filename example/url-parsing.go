package main

import (
	"net/url"
	"fmt"
	"net"
)

func main() {
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	fmt.Println(u)
	fmt.Println(u.Scheme)
	fmt.Println(u.Host)
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	fmt.Println(u.User.Password())

	host,port,_ := net.SplitHostPort(u.Host)
	fmt.Println(host, port)
}