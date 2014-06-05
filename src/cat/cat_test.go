package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func ExampleCatWithOneFile() {
	var files = []string{

		filepath.Join("..", "..", "resources", "test_001.txt"),
	}
	Cat(files, os.Stdout)
	// Output: hello
}

func ExampleCatWithTwoFiles() {
	var files = []string{
		filepath.Join("..", "..", "resources", "test_001.txt"),
		filepath.Join("..", "..", "resources", "test_002.txt"),
	}
	Cat(files, os.Stdout)
	// Output: helloworld
}

func ExampleCatFileWithOneFile() {
	fd, err := os.Open(filepath.Join("..", "..", "resources", "test_001.txt"))
	if err != nil {
		log.Fatal("fail - file not found\n")
	}
	defer fd.Close()
	CatFile(fd, os.Stdout)
	// Output: hello
}

func BenchmarkCatFile001(b *testing.B) {
	for n:= 0; n < b.N; n++ {
		fd, err := os.Open(filepath.Join("..", "..", "resources", "cat_benchmark.txt"))
		if err != nil {
			log.Fatal("fail - file not found: " + err.Error())
		}
		CatFile(fd, ioutil.Discard)
		fd.Close() // explicit to avoid opening too many files
	}
}
