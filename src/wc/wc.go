package main

import (
	"flag"
	"fmt"
	"os"
	"gobase"
)

var (
	countLines = flag.Bool("l", false, "count lines")
	countWords = flag.Bool("w", false, "count words")
	countBytes = flag.Bool("c", false, "count bytes")
	countChars = flag.Bool("m", false, "count characters")
)

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		wc(os.Stdin)
	}
	for i := 0; i < flag.NArg(); i++ {
		fd, er := os.Open(flag.Arg(i))
		defer fd.Close()
		if er != nil {
			fmt.Fprintln(os.Stderr, "wc: ", er)
		} else {
			wc(fd)
		}
	}
}

func wc(fd *os.File) {
	wc, ew := gobase.Wc(fd)
	if ew != nil {
		fmt.Fprintln(os.Stderr, "wc: ", ew)
	} else {
		output(wc, fd.Name())
	}
}

func output(wc gobase.WordCount, name string) {
	noFlags := !*countLines && !*countWords && !*countBytes && !*countChars
	
	if *countLines || noFlags {
		fmt.Printf(" %5d", wc.Lines)
	}
	if *countWords || noFlags {
		fmt.Printf(" %5d", wc.Words)
	}
	if *countBytes || noFlags {
		fmt.Printf(" %5d", wc.Bytes)
	}
	if *countChars {
		fmt.Printf(" %5d", wc.Chars)
	}
	fmt.Printf(" %s\n", name)
}
