package main

import (
	"flag"
	"fmt"
	"os"
	. "gobase"
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
    fmt.Println(Basename(flag.Arg(0)))
}
