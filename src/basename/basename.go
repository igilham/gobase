package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"
)

// strip directory from filenames
func main() {
	flag.Parse()
	switch {
	case flag.NArg() == 0:
		log.Fatal("basename: not enough arguments")
	case flag.NArg() > 1:
		log.Fatal("basename: too many arguments")
	}
	fmt.Println(filepath.Base(flag.Arg(0)))
}
