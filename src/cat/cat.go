package main

import (
	"errors"
	"flag"
	"log"
	"os"
)

const (
	nbuf = 4096 // buffer size in bytes
)

// Cat concatenates files.
// any errors encountered will stop iteration through the list of files.
func Cat(files []string) error {
	for _, s := range files {
		var fd *os.File
		var er error
		if s == "-" {
			fd = os.Stdin
		} else {
			fd, er = os.Open(s)
			defer fd.Close()
			if er != nil {
				return er
			}
		}
		ew := CatFile(fd)
		if ew != nil {
			return er
		}
	}
	return nil
}

// CatFile reads a single file and writes to os.Stdout.
func CatFile(fd *os.File) error {
	var buf [nbuf]byte
	for {
		switch nr, er := fd.Read(buf[:]); true {
		case nr < 0:
			return errors.New("error reading from " + fd.Name() + " : " + er.Error())
		case nr == 0: // EOF
			return nil
		case nr > 0:
			if nw, _ := os.Stdout.Write(buf[0:nr]); nw != nr {
				return errors.New("error reading from " + fd.Name() + " : " + er.Error())
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
		files = append(files, "-")
	}
	for i := 0; i < flag.NArg(); i++ {
		files = append(files, flag.Arg(i))
	}

	handleError(Cat(files))
}

// handle errors in the cat process
func handleError(er error) {
	if er != nil {
		log.Fatalf("cat: %v", er)
	}
}
