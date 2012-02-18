package main

import (
	"flag"
	"fmt"
	"../gobase"
	"os"
)

// concatenate the specified files, joining with a newline. If -n is specified, 
// the newline will be omitted.
func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		handleError(gobase.Cat(os.Stdin))
	}
	for i := 0; i < flag.NArg(); i++ {
		fd, err := os.Open(flag.Arg(i))
		defer fd.Close()
		if err != nil {
			fmt.Fprintln(os.Stderr, fmt.Sprintf("cat: can't open %s: %s", flag.Arg(i), err))
			os.Exit(1)
		}
		handleError(gobase.Cat(fd))
	}
}

func handleError(er error) {
	if er != nil {
		fmt.Fprintln(os.Stderr, "cat: ", er)
	}
}
