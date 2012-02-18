package main

import (
	"flag"
	"fmt"
	"../gobase"
	"os"
	"sort"
)

var reverse = flag.Bool("r", false, "reverse sort order")

// sort input lines
func main() {
	flag.Parse()
	var lines []string
	if flag.NArg() == 0 {
		lines = gobase.Head(os.Stdin, 0)
	} else {
		for i := 0; i < flag.NArg(); i++ {
			file, err := os.Open(flag.Arg(i))
			if err != nil {
				fmt.Fprintln(os.Stderr, "sort: cannot open file ", flag.Arg(i))
			} else {
				lines = gobase.Head(file, 0)
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
