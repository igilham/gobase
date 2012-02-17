package main

import (
	"flag"
	"fmt"
	"os"
	. "gobase"
)

var lines = flag.Int("n", 10, "number of lines to print")

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		head(os.Stdin, *lines)
	} else {
		for i, arg := range flag.Args() {
			if f, err := os.Open(arg); err != nil {
				fmt.Fprintln(os.Stderr, "head: cannot open file %s", arg)
			} else {
				head(f, *lines)
			}
			if i < flag.NArg() - 1 {
				fmt.Println()
			}
		}
	}
}

func head(f *os.File, n int) {
	for _, s := range Head(f, n) {
		fmt.Println(s)
	}
}
