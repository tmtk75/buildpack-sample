package main

import (
	"flag"
	"fmt"
)

var Commit string

func main() {
	msg := flag.String("msg", "", "message")
	flag.Parse()

	fmt.Println("commit:", Commit)
	fmt.Println("msg:", *msg)
}
