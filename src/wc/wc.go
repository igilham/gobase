package main

import (
	"flag"
	"fmt"
	"os"
	"gobase"
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
		fmt.Println(fmt.Sprintf("  %d  %d  %d  %s", wc.Lines, wc.Words, wc.Bytes, fd.Name()))
	}
}
