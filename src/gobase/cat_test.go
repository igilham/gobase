package gobase

import (
	"fmt"
	"os"
)

func ExampleCat_withOneFile() {
	var files = []string{
		"test_001.txt",
	}
	Cat(files)
	// Output: hello
}

func ExampleCat_withTwoFiles() {
	var files = []string{
		"test_001.txt",
		"test_002.txt",
	}
	Cat(files)
	// Output: helloworld
}

func ExampleCatFile_withOneFile() {
	fd, err := os.Open("test_001.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "fail - file not found\n")
		os.Exit(1)
	}
	defer fd.Close()
	CatFile(fd)
	// Output: hello
}
