package main

import (
	"flag"
	"fmt"
	"os"
)

// Assuming UMASK is 022 (0644 for files and 0755 for folders)
var mode = flag.Uint("m", uint(0755), "Set the mode for the folder")

func main() {
	flag.Parse()
	for i := 0; i < flag.NArg(); i++ {
		m := os.FileMode(*mode)
		err := os.Mkdir(flag.Arg(i), m)
		if err != nil {
			fmt.Fprintf(os.Stderr, "mkdir: %v\n", err)
		}
	}
}
