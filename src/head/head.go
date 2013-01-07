package main

import (
	"../gobase"
	"flag"
	"fmt"
	"os"
)

var lines = flag.Int("n", 10, "number of lines to print")

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		handle(gobase.Head(os.Stdin, *lines))
	} else {
		for i, arg := range flag.Args() {
			if f, err := os.Open(arg); err != nil {
				fmt.Fprintln(os.Stderr, "head: cannot open file ", arg)
				os.Exit(1)
			} else {
				defer f.Close()
				handle(gobase.Head(f, *lines))
				if i < flag.NArg()-1 {
					fmt.Println()
				}
			}
		}
	}
}

func handle(er error) {
	if er != nil {
		fmt.Fprintf(os.Stderr, "head: %v\n", er)
	}
}
