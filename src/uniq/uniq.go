package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"os"
)

// Uniq prints the input, omitting duplicated lines
func Uniq(fd *os.File) error {
	var prevLine string
	reader := bufio.NewReader(fd)
	buffer := new(bytes.Buffer)
	firstLine := true
	for {
		part, prefix, er := reader.ReadLine()
		if er != nil {
			break
		}
		buffer.Write(part)
		if !prefix {
			thisLine := buffer.String()
			if firstLine || prevLine != thisLine {
				fmt.Println(thisLine)
				firstLine = false
			}
			prevLine = thisLine
			buffer.Reset()
		}
	}
	return nil
}

// print unique lines from input
func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		handleError(Uniq(os.Stdin))
	}
	for i := 0; i < flag.NArg(); i++ {
		fd, err := os.Open(flag.Arg(i))
		handleError(err)
		defer fd.Close()
		handleError(Uniq(fd))
	}
}

func handleError(er error) {
	if er != nil {
		fmt.Fprintln(os.Stderr, "uniq: %v", er)
		os.Exit(1)
	}
}
