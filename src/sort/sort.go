package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"os"
	"sort"
)

var reverse = flag.Bool("r", false, "reverse sort order")

var lines []string

// sort input lines
func main() {
	flag.Parse()
    if flag.NArg() == 0 {
		getlines(os.Stdin)
    } else {
		for i := 0; i < flag.NArg(); i++ {
			if file, err := os.Open(flag.Arg(i)); err != nil {
				fmt.Fprintln(os.Stderr, "sort: cannot open file %s",
					flag.Arg(i))
			} else {
				getlines(file)
			}
		}
    }
    if len(lines) > 0 {
		sort.Strings(lines)
		out(lines)
	}
}

func getlines(file *os.File) {
	var part []byte
	var	prefix bool
	var err os.Error
	reader := bufio.NewReader(file)
	buffer := bytes.NewBuffer(make([]byte, 1024))
	for {
		if part, prefix, err = reader.ReadLine(); err != nil {
			break
		}
		buffer.Write(part)
		if !prefix {
			lines = append(lines, buffer.String())
			buffer.Reset()
		}
	}
}

func out(list []string) {
	var (
		start int64 = 0
		end int64 = int64(len(list))
		step int64 = 1
	)
	if *reverse {
		step = -1
		start = end - 1
		end = -1
	}
	for i := start; i != end; i += step {
		fmt.Println(list[i])
	}
}
