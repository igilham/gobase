package main

import (
	"flag"
	"fmt"
	"strings"
)

var output string

func main() {
	output = whatToPrint()
	for {
		fmt.Println(output)
	}
}

func whatToPrint() (what string) {
	flag.Parse()
	if flag.NArg() > 0 {
		what = strings.Join(flag.Args(), " ")
	} else {
		what = "y"
	}
	return
}
