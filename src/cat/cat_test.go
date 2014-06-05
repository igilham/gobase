package main

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func ExampleCatWithOneFile() {
	var files = []string{
		"../../resources/test_001.txt",
	}
	Cat(files, os.Stdout)
	// Output: hello
}

func ExampleCatWithTwoFiles() {
	var files = []string{
		"../../resources/test_001.txt",
		"../../resources/test_002.txt",
	}
	Cat(files, os.Stdout)
	// Output: helloworld
}

func ExampleCatFileWithOneFile() {
	fd, err := os.Open("../../resources/test_001.txt")
	if err != nil {
		log.Fatal("fail - file not found\n")
	}
	defer fd.Close()
	CatFile(fd, os.Stdout)
	// Output: hello
}

func BenchmarkCatFile001(b *testing.B) {
	for n:= 0; n < b.N; n++ {
		fd, err := os.Open("../../resources/cat_benchmark.txt")
		if err != nil {
			log.Fatal("fail - file not found: " + err.Error())
		}
		CatFile(fd, ioutil.Discard)
		fd.Close() // explicit to avoid opening too many files
	}
}
