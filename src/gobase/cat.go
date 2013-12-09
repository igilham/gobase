package gobase

import (
	"errors"
	"os"
)

// Cat concatenates files.
// any errors encountered will stop iteration through the list of files.
func Cat(files []string) error {
	for _, s := range files {
		fd, er := os.Open(s)
		defer fd.Close()
		if er != nil {
			return er
		}
		ew := CatFile(fd)
		if ew != nil {
			return ew
		}
	}
	return nil
}

// CatFile reads a single file and writes to os.Stdout.
func CatFile(fd *os.File) error {
	const NBUF = 512
	var buf [NBUF]byte
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
