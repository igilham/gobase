package gobase

import (
	"errors"
	"os"
)

// Cat concatenates files.
func Cat(fd *os.File) error {
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
	return nil
}
