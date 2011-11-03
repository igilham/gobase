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

// sort input lines
func main() {
	flag.Parse()
	var lines []string
    if flag.NArg() == 0 {
		lines = getlines(os.Stdin)
    } else {
		for i := 0; i < flag.NArg(); i++ {
			file, err := os.Open(flag.Arg(i))
			if err != nil {
				fmt.Fprintln(os.Stderr, "sort: cannot open file %s", flag.Arg(i))
			} else {
				lines = getlines(file)
			}
		}
    }
    if len(lines) > 0 {
		sort.Strings(lines)
		out(lines)
	}
}

func getlines(file *os.File) ([]string) {
	var lines []string
	var part []byte
	var	prefix bool
	var err os.Error
	reader := bufio.NewReader(file)
	buffer := bytes.NewBuffer(make([]byte, 4096))
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
	return lines
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
