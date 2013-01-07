package main

import (
	"../gobase"
	"flag"
	"log"
	"os"
)

// print unique lines from input
func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		handleError(gobase.Uniq(os.Stdin))
	}
	for i := 0; i < flag.NArg(); i++ {
		fd, err := os.Open(flag.Arg(i))
		handleError(err)
		defer fd.Close()
		handleError(gobase.Uniq(fd))
	}
}

func handleError(er error) {
	if er != nil {
		log.Fatalf("uniq: %v", er)
	}
}
