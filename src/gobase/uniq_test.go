package gobase

import (
	"fmt"
	"log"
	"os"
)

func ExampleUniq() {
	fd, er := os.Open("../../resources/test_005.txt")
	if er != nil {
		log.Fatalf("%v\n", er)
	}
	defer fd.Close()
	ew := Uniq(fd)
	if ew != nil {
		fmt.Fprintf(os.Stderr, "%v\n", ew)
	}
	// Output:
	// hello
	// world
	// hello
}
