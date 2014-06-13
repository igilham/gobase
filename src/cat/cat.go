package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

const (
	// buffer size in bytes
	// benchmarking has shown that 4kB is a good balance
	nbuf             = 1024 * 4
	stdinPlaceholder = "-"
)

// Cat concatenates files.
// any errors encountered will stop iteration through the list of files.
// "-" counts as os.Stdin
func Cat(files []string, writer io.Writer) error {
	for _, s := range files {
		var fd *os.File
		var er error
		if s == stdinPlaceholder {
			fd = os.Stdin
		} else {
			fd, er = os.Open(s)
			defer fd.Close()
			if er != nil {
				return er
			}
		}
		ew := CatFile(fd, writer)
		if ew != nil {
			return er
		}
	}
	return nil
}

// CatFile reads a single file and writes to os.Stdout.
func CatFile(reader io.Reader, writer io.Writer) error {
	var buf [nbuf]byte
	for {
		nr, er := reader.Read(buf[:])
		switch {
		case nr < 0:
			return errors.New("read error: " + er.Error())
		case nr == 0: // EOF
			return nil
		case nr > 0:
			if nw, ew := writer.Write(buf[0:nr]); nw != nr {
				return errors.New("write error: " + ew.Error())
			}
		}
	}
}

// concatenate the specified files, joining with a newline. If -n is specified,
// the newline will be omitted.
func main() {
	flag.Parse()
	var files []string
	if flag.NArg() == 0 {
		files = append(files, stdinPlaceholder)
	}
	for i := 0; i < flag.NArg(); i++ {
		files = append(files, flag.Arg(i))
	}

	handleError(Cat(files, os.Stdout))
}

// handle errors in the cat process
func handleError(er error) {
	if er != nil {
		fmt.Fprintf(os.Stderr, "cat: %v\n", er)
	}
}
