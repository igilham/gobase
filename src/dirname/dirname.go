package main

import (
    "flag"
    "fmt"
    "os"
    "strings"
)

const (
    cwd string = "."
    sep = string(os.PathSeparator)
)

// strip non-directory suffix from filename
func main() {
    flag.Parse()
    switch {
        case flag.NArg() == 0:
            fmt.Fprintf(os.Stderr, "dirname: not enough arguments")
            os.Exit(1)
        case flag.NArg() > 1:
            fmt.Fprintf(os.Stderr, "dirname: too many arguments")
            os.Exit(1)
    }
    fmt.Println(dirname(flag.Arg(0)))
}

func dirname (s string) string {
    if s == sep {
        return s
    }
    s = strings.TrimRight(s, sep)
    index := strings.LastIndex(s, sep)
    if index == -1 {
        return cwd
    }
    // remove anything after trailing /
    // cast to []int avoids multibyte unicode issues
    s = string([]int(s)[0:index+1])
    if len(s) <= 1 {
		return s
    }
    return strings.TrimRight(s, sep)
}
