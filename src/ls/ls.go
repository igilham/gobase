package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

const (
	colon = ":"
)

// list files - lists files in the current directory if no argument is given, 
// or lists files in all directories specified as arguments
func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		list(".")
	} else {
		for i := 0; i < flag.NArg(); i++ {
			fmt.Println(flag.Arg(i) + colon)
			list(flag.Arg(i))
			if i < flag.NArg()-1 {
				fmt.Println()
			}
		}
	}
}

func list(s string) {
	file, err := os.Open(s)
	if err != nil {
		fmt.Fprintln(os.Stderr, "ls: error accessing ", s)
		os.Exit(1)
	}
	defer file.Close()
	subfiles, err2 := file.Readdirnames(0)
	if len(subfiles) == 0 && err2 != nil {
		log.Fatalf("ls: error accessing contents of ", s)
	}
	for i := 0; i < len(subfiles); i++ {
		fmt.Println(subfiles[i])
	}
}
