package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"os"
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
	for _, s := range getlines(f, n) {
		fmt.Println(s)
	}
}

func getlines(file *os.File, n int) ([]string) {
	var lines []string
	var part []byte
	var	prefix bool
	var err os.Error
	reader := bufio.NewReader(file)
	buffer := bytes.NewBuffer(make([]byte, 4096))
	for n != len(lines) {
		if part, prefix, err = reader.ReadLine(); err != nil {
			break
		}
		buffer.Write(part)
		if !prefix {
			lines = append(lines, buffer.String())
			buffer.Reset()
		}
	}
	return lines
}
