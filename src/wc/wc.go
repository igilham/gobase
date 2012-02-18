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
	oneFlag := !(*countLines && *countWords) &&
		!(*countLines && *countBytes) &&
		!(*countLines && *countChars) &&
		!(*countWords && *countBytes) &&
		!(*countWords && *countChars) &&
		(*countLines || *countWords || *countBytes || *countChars)
	format := " %5d"
	if (oneFlag) {
		format = "%d"
	}
	if *countLines || noFlags {
		fmt.Printf(format, wc.Lines)
	}
	if *countWords || noFlags {
		fmt.Printf(format, wc.Words)
	}
	if *countBytes || noFlags {
		fmt.Printf(format, wc.Bytes)
	}
	if *countChars {
		fmt.Printf(format, wc.Chars)
	}
	fmt.Printf(" %s\n", name)
}
