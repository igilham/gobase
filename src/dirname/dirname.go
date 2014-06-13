package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

const (
	// string representing the current path in the operating system
	cwd string = "."
	// file path separator
	separator = string(os.PathSeparator)
)

// Dirname strips the filename from a directory path.
func Dirname(s string) string {
	switch {
		case s == "":
			return "."
		case s == separator:
			return s
	}
	hasSlashes := strings.Contains(s, separator)
	s = strings.TrimRight(s, separator)
	if hasSlashes && s == "" {
		return separator
	}
	index := strings.LastIndex(s, separator)
	if index == -1 {
		return cwd
	}
	// remove anything after trailing /
	// cast to []rune avoids multibyte unicode issues
	s = string([]rune(s)[0 : index+1])
	if len(s) <= 1 {
		return s
	}
	return strings.TrimRight(s, separator)
}

// strip non-directory suffix from filename
func main() {
	flag.Parse()
	if flag.NArg() == 0 || flag.NArg() > 1 {
		fmt.Fprintf(os.Stderr, "usage: dirname path")
		os.Exit(1)
	} else {
		fmt.Println(Dirname(flag.Arg(0)))
	}
}
