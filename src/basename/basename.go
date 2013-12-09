package main

import (
	"flag"
	"fmt"
	"gobase"
	"log"
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
	fmt.Println(gobase.Basename(flag.Arg(0)))
}
