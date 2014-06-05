package main

import (
	"log"
	"os"
	"testing"
)

func ExampleCatWithOneFile() {
	var files = []string{
		"../../resources/test_001.txt",
	}
	Cat(files)
	// Output: hello
}

func ExampleCatWithTwoFiles() {
	var files = []string{
		"../../resources/test_001.txt",
		"../../resources/test_002.txt",
	}
	Cat(files)
	// Output: helloworld
}

func ExampleCatFileWithOneFile() {
	fd, err := os.Open("../../resources/test_001.txt")
	if err != nil {
		log.Fatal("fail - file not found\n")
	}
	defer fd.Close()
	CatFile(fd)
	// Output: hello
}

func BenchmarkCatFileWithOneLargeFile(b *testing.B) {
	fd, err := os.Open("../../resources/test_cat_large.txt")
	if err != nil {
		log.Fatal("fail - file not found\n")
	}
	defer fd.Close()
	CatFile(fd)
}
