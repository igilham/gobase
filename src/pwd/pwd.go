package main

import (
	"fmt"
	"log"
	"os"
)

// prints the current working directory
func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("pwd: %v", err)
	} else {
		fmt.Println(wd)
	}
}
