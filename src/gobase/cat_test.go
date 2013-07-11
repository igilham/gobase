package gobase

import (
	"log"
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
		log.Fatal("fail - file not found\n")
	}
	defer fd.Close()
	CatFile(fd)
	// Output: hello
}
