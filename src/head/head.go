package main

import (
	"gobase"
	"flag"
	"fmt"
	"log"
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
				log.Fatalf("head: cannot open file \n", arg)
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
