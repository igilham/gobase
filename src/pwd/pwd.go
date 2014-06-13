package main

import (
	"fmt"
	"os"
)

// prints the current working directory
func main() {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "pwd: %v\n", err)
		os.Exit(1)
	} else {
		fmt.Println(wd)
	}
}
