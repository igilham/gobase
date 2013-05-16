package main

import (
	"gobase"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
)

var reverse = flag.Bool("r", false, "reverse sort order")

// sort input lines
func main() {
	flag.Parse()
	var lines []string
	var er error
	if flag.NArg() == 0 {
		lines, er = gobase.GetLines(os.Stdin, 0)
		if er != nil {
			fmt.Fprintf(os.Stderr, "sort: %v", er)
		}
	} else {
		for i := 0; i < flag.NArg(); i++ {
			file, err := os.Open(flag.Arg(i))
			if err != nil {
				log.Fatalf("sort: cannot open file ", flag.Arg(i))
			}
			defer file.Close()
			var ew error
			if lines, ew = gobase.GetLines(file, 0); ew != nil {
				fmt.Fprintf(os.Stderr, "sort: %v", ew)
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
