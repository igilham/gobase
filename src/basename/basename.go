package main

import (
	"flag"
	"fmt"
	"../gobase"
	"os"
)

// strip directory from filenames
func main() {
	flag.Parse()
	switch {
	case flag.NArg() == 0:
		fmt.Fprintln(os.Stderr, "basename: not enough arguments")
		os.Exit(1)
	case flag.NArg() > 1:
		fmt.Fprintln(os.Stderr, "basename: too many arguments")
		os.Exit(1)
	}
	fmt.Println(gobase.Basename(flag.Arg(0)))
}
