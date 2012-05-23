package main

import (
	"flag"
	"fmt"
	"../gobase"
	"os"
)

// print unique lines from input
func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		handleError(gobase.Uniq(os.Stdin))
	}
	for i := 0; i < flag.NArg(); i++ {
		fd, err := os.Open(flag.Arg(i))
		defer fd.Close()
		handleError(err)
		handleError(gobase.Uniq(fd))
	}
}

func handleError(er error) {
	if er != nil {
		fmt.Fprintln(os.Stderr, "uniq: ", er)
		os.Exit(1)
	}
}