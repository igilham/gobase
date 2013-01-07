package gobase

import (
	"fmt"
	"os"
)

func ExampleUniq() {
	fd, er := os.Open("test_005.txt")
	if er != nil {
		fmt.Fprintf(os.Stderr, "%s\n", er)
		os.Exit(1)
	}
	defer fd.Close()
	ew := Uniq(fd)
	if ew != nil {
		fmt.Fprintf(os.Stderr, "%s\n", ew)
	}
	// Output:
	// hello
	// world
	// hello
}
