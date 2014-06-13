package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"os"
)

var lines = flag.Int("n", 10, "number of lines to print")

// Head gets the first n lines of a file. If n == 0, then attempt to get all
// the lines.
func Head(file *os.File, n int) error {
	lines := 0
	var part []byte
	var prefix bool
	var err error
	reader := bufio.NewReader(file)
	buffer := new(bytes.Buffer)
	for n != lines || n == 0 {
		if part, prefix, err = reader.ReadLine(); err != nil {
			return err
		}
		buffer.Write(part)
		if !prefix {
			fmt.Printf("%s\n", buffer)
			buffer.Reset()
			lines++
		}
	}
	return nil
}

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		handle(Head(os.Stdin, *lines))
	} else {
		for i, arg := range flag.Args() {
			if f, err := os.Open(arg); err != nil {
				fmt.Fprintf(os.Stderr, "head: cannot open file %s\n", arg)
				os.Exit(1)
			} else {
				defer f.Close()
				handle(Head(f, *lines))
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
