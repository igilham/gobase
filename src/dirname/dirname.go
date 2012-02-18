package main

import (
	"flag"
	"fmt"
	"../gobase"
	"os"
)

// strip non-directory suffix from filename
func main() {
	flag.Parse()
	switch {
	case flag.NArg() == 0:
		fmt.Fprintln(os.Stderr, "dirname: not enough arguments")
		os.Exit(1)
	case flag.NArg() > 1:
		fmt.Fprintln(os.Stderr, "dirname: too many arguments")
		os.Exit(1)
	}
	fmt.Println(gobase.Dirname(flag.Arg(0)))
}
