package main

import (
	"fmt"
	"os"
)

// prints the current working directory
func main() {
	wd, err := os.Getwd(); if err != nil {
		fmt.Fprintln(os.Stderr, "pwd: ", err)
		os.Exit(1)
	} else {
		fmt.Println(wd)
	}
}
