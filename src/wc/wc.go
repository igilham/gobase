package main

import (
	"flag"
	"fmt"
	"os"
	"gobase"
)

func main() {
	flag.Parse()
	for i := 0; i < flag.NArg(); i++ {
		fd, er := os.Open(flag.Arg(i))
		defer fd.Close()
		if er != nil {
			fmt.Fprintln(os.Stderr, "wc: ", er)
		} else {
			wc, ew := gobase.Wc(fd)
			if ew != nil {
				fmt.Fprintln(os.Stderr, "wc: ", er)
			} else {
				fmt.Println(fmt.Sprintf("  %d  %d  %d  %s", wc.Lines, wc.Words, wc.Bytes, flag.Arg(i)))
			}
		}
	}
}
