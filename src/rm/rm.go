package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		fmt.Fprintf(os.Stderr, "rm: no file specified")
		os.Exit(1)
	}
	exitCode := 0
	for i := 0; i < flag.NArg(); i++ {
		er := rm(flag.Arg(i))
		if er != nil {
			fmt.Fprintf(os.Stderr, "rm: %s", er)
			exitCode = 1
		}
	}
	os.Exit(exitCode)
}

func rm(path string) os.Error {
	info, er := os.Lstat(path)
	switch {
		case er != nil:
			return er
		case info.IsDirectory():
			return os.NewError(path + " is a directory")
		default:
			return os.Remove(path)
	}
	return nil
}
