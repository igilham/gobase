package main

import (
	"log"
	"os"
)

func ExampleHead_1Line() {
	fd, er := os.Open("../../resources/test_004.txt")
	if er != nil {
		log.Fatalf("%v\n", er)
	}
	defer fd.Close()
	Head(fd, 1)
	// Output: 001
}

func ExampleHead_AllLinesInOneLineFile() {
	fd, er := os.Open("../../resources/test_001.txt")
	if er != nil {
		log.Fatalf("%v\n", er)
	}
	defer fd.Close()
	Head(fd, 0)
	// Output: hello
}

func ExampleHead_10Lines() {
	fd, er := os.Open("../../resources/test_004.txt")
	if er != nil {
		log.Fatalf("%v\n", er)
	}
	defer fd.Close()
	Head(fd, 10)
	// Output:
	// 001
	// 002
	// 003
	// 004
	// 005
	// 006
	// 007
	// 008
	// 009
	// 010
}

func ExampleHead_AllLines() {
	fd, er := os.Open("../../resources/test_004.txt")
	if er != nil {
		log.Fatalf("%v\n", er)
	}
	defer fd.Close()
	Head(fd, 0)
	// Output:
	// 001
	// 002
	// 003
	// 004
	// 005
	// 006
	// 007
	// 008
	// 009
	// 010
	// 011
	// 012
	// 013
	// 014
	// 015
	// 016
	// 017
	// 018
	// 019
	// 020
}

func ExampleHead_19Lines() {
	fd, er := os.Open("../../resources/test_004.txt")
	if er != nil {
		log.Fatalf("%v\n", er)
	}
	defer fd.Close()
	Head(fd, 19)
	// Output:
	// 001
	// 002
	// 003
	// 004
	// 005
	// 006
	// 007
	// 008
	// 009
	// 010
	// 011
	// 012
	// 013
	// 014
	// 015
	// 016
	// 017
	// 018
	// 019
}
