package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	// file path separator
	sep = string(os.PathSeparator)
)

// Basename strips the non-directory suffix from a filename.
func Basename(s string) string {
	arr := strings.Split(s, string(os.PathSeparator))
	for i := len(arr) - 1; i >= 0; i-- {
		if len(arr[i]) != 0 {
			s = arr[i]
			break
		}
		if i == 0 {
			return string(os.PathSeparator)
		}
	}
	return s
}

// strip directory from filenames
func main() {
	flag.Parse()
	switch {
	case flag.NArg() == 0:
		log.Fatal("basename: not enough arguments")
	case flag.NArg() > 1:
		log.Fatal("basename: too many arguments")
	}
	fmt.Println(Basename(flag.Arg(0)))
}
