package main

import (
	"os"
	"flag"
	"fmt"
)

// Assuming UMASK is 022 (0644 for files and 0755 for folders)
var mode = flag.Uint("m", 0755, "Set the mode for the folder")

func main() {
	flag.Parse()
	for i := 0; i < flag.NArg(); i++ {
		err := os.Mkdir(flag.Arg(i), uint32(*mode))
		if err != nil { fmt.Fprintln(os.Stderr, err) }
	}
}