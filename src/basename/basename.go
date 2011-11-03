package main

import (
	"flag"
	"fmt"
	"os"
    "strings"
)

const sep = string(os.PathSeparator)

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
    fmt.Println(basename(flag.Arg(0)))
}

func basename(s string) string {
	arr := strings.Split(s, sep)
	for i := len(arr) - 1; i >= 0; i-- {
		if len(arr[i]) != 0 {
			s = arr[i]
			break
		}
		if i == 0 {
			return sep
		}
	}
	return s
}
