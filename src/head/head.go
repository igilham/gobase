package main

import (
	"flag"
	"fmt"
	"../gobase"
	"os"
)

var lines = flag.Int("n", 10, "number of lines to print")

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		head(os.Stdin, *lines)
	} else {
		for i, arg := range flag.Args() {
			f, err := os.Open(arg)
			defer f.Close()
			if err != nil {
				fmt.Fprintln(os.Stderr, "head: cannot open file ", arg)
			} else {
				head(f, *lines)
			}
			if i < flag.NArg()-1 {
				fmt.Println()
			}
		}
	}
}

func head(f *os.File, n int) {
	for _, s := range gobase.Head(f, n) {
		fmt.Println(s)
	}
}
