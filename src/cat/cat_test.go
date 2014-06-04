package main

import (
	"log"
	"os"
)

func ExampleCat_withOneFile() {
	var files = []string{
		"../../resources/test_001.txt",
	}
	Cat(files)
	// Output: hello
}

func ExampleCat_withTwoFiles() {
	var files = []string{
		"../../resources/test_001.txt",
		"../../resources/test_002.txt",
	}
	Cat(files)
	// Output: helloworld
}

func ExampleCatFile_withOneFile() {
	fd, err := os.Open("../../resources/test_001.txt")
	if err != nil {
		log.Fatal("fail - file not found\n")
	}
	defer fd.Close()
	CatFile(fd)
	// Output: hello
}
