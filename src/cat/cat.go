package main

import (
	"../gobase"
	"flag"
	"log"
	"os"
)

// concatenate the specified files, joining with a newline. If -n is specified, 
// the newline will be omitted.
func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		handleError(gobase.CatFile(os.Stdin))
	}
	var files []string
	for i := 0; i < flag.NArg(); i++ {
		files = append(files, flag.Arg(i))
	}

	handleError(gobase.Cat(files))
}

// handle errors in the cat process
func handleError(er error) {
	if er != nil {
		log.Fatalf("cat: %v", er)
	}
}
