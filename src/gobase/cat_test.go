package gobase

import (
	"fmt"
	"os"
)

func ExampleCatWithOneFile() {
	var files = []string {
		"cat_test_001.txt",
	}
	Cat(files);
	// Output: hello
}

func ExampleCatWithTwoFiles() {
	var files = []string {
		"cat_test_001.txt",
		"cat_test_002.txt",
	}
	Cat(files);
	// Output: helloworld
}

func ExampleCatFileWithOneFile() {
	fd, err := os.Open("cat_test_001.txt")
	defer fd.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fail - file not found\n")
	} else {
		CatFile(fd);
	}
	// Output: hello
}
