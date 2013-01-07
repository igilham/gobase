package main

import (
	"../gobase"
	"flag"
	"fmt"
	"log"
)

// strip non-directory suffix from filename
func main() {
	flag.Parse()
	switch {
	case flag.NArg() == 0:
		log.Fatalf("dirname: not enough arguments")
	case flag.NArg() > 1:
		log.Fatalf("dirname: too many arguments")
	}
	fmt.Println(gobase.Dirname(flag.Arg(0)))
}
