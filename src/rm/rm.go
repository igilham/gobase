package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		fmt.Fprintln(os.Stderr, "usage: rm FILES")
		os.Exit(1)
	}
	exitCode := 0
	for i := 0; i < flag.NArg(); i++ {
		er := rm(flag.Arg(i))
		if er != nil {
			fmt.Fprintf(os.Stderr, "rm: %v\n", er)
			exitCode = 1
		}
	}
	os.Exit(exitCode)
}

func rm(path string) error {
	info, er := os.Lstat(path)
	switch {
	case er != nil:
		return er
	case info.IsDir():
		return errors.New(path + " is a directory")
	default:
		return os.Remove(path)
	}
}
