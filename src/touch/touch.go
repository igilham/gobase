package main

import (
	"flag"
	"fmt"
	"gobase"
	"log"
	"os"
)

// Change file timestamps or create files that don't already exist
func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		log.Fatal("touch: missing file operand")
	}
	for i := 0; i < flag.NArg(); i++ {
		er := gobase.Touch(flag.Arg(i))
		if er != nil {
			fmt.Fprintln(os.Stderr, "touch: cannot create file ", flag.Arg(i))
		}
	}
}
