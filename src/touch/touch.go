package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

// Touch updates the access and modify times of a file, or creates a new
// file if it doesn't already exist.
func Touch(path string) error {
	now := time.Now()
	er := os.Chtimes(path, now, now)
	if er != nil {
		_, ew := os.Create(path)
		return ew
	}
	return nil
}

// Change file timestamps or create files that don't already exist
func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		fmt.Fprintln(os.Stderr, "usage: touch path...")
		os.Exit(1)
	}
	for i := 0; i < flag.NArg(); i++ {
		er := Touch(flag.Arg(i))
		if er != nil {
			fmt.Fprintf(os.Stderr, "touch: cannot create file %s\n", flag.Arg(i))
		}
	}
}
