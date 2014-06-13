package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

// strip directory from filenames
func main() {
	flag.Parse()
	if flag.NArg() == 0 || flag.NArg() > 1 {
		fmt.Fprintln(os.Stderr, "usage: basename string")
		os.Exit(1)
	} else {
		fmt.Println(filepath.Base(flag.Arg(0)))
	}
}
