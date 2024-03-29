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

// GetLines reads the first n lines of the file f.
// The end result is similar to Head, but suited to different use cases.
// It returns a slice of strings for each line read and an error if encountered.
func GetLines(file *os.File, n int) ([]string, error) {
	var lines []string
	var part []byte
	var prefix bool
	var err error
	reader := bufio.NewReader(file)
	buffer := new(bytes.Buffer)
	for n != len(lines) || n == 0 {
		if part, prefix, err = reader.ReadLine(); err != nil {
			return lines, err
		}
		buffer.Write(part)
		if !prefix {
			lines = append(lines, buffer.String())
			buffer.Reset()
		}
	}
	return lines, nil
}

// sort input lines
func main() {
	flag.Parse()
	var lines []string
	var er error
	if flag.NArg() == 0 {
		lines, er = GetLines(os.Stdin, 0)
		if er != nil {
			fmt.Fprintf(os.Stderr, "sort: %v\n", er)
		}
	} else {
		for i := 0; i < flag.NArg(); i++ {
			file, err := os.Open(flag.Arg(i))
			if err != nil {
				fmt.Fprintf(os.Stderr, "sort: cannot open file %s\n", flag.Arg(i))
				os.Exit(1)
			}
			defer file.Close()
			var ew error
			if lines, ew = GetLines(file, 0); ew != nil {
				fmt.Fprintf(os.Stderr, "sort: %v\n", ew)
			}
		}
	}
	if len(lines) > 0 {
		sort.Strings(lines)
		out(lines)
	}
}

func out(list []string) {
	var (
		start int64 = 0
		end   int64 = int64(len(list))
		step  int64 = 1
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
