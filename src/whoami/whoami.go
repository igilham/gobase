package main

import (
	"flag"
	"fmt"
	"os"
	"os/user"
)

// Prints the effective user ID.
func main() {
	flag.Parse()
	user, error := user.Current()
	if error != nil {
		fmt.Fprintf(os.Stderr, "whoami: %v\n", error)
		os.Exit(1)
	}
	fmt.Println(user.Username)
}
