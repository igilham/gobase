package gobase

import (
	"errors"
	"hash"
	"hash/crc32"
	"os"
)

// Cksum calculates a simple checksum for a file. Currently uses crc32.
func Cksum(fd *os.File) (hash.Hash32, error) {
	table := crc32.MakeTable(crc32.Castagnoli)
	h := crc32.New(table)
	const NBUF = 512
	var buf [NBUF]byte
	for {
		switch nr, er := fd.Read(buf[:]); true {
		case nr < 0:
			return nil, errors.New("error reading from " + fd.Name() + " : " + er.Error())
		case nr == 0: // EOF
			break
		case nr > 0:
			_, ew := h.Write(buf[0:nr])
			if ew != nil {
				return nil, errors.New("error writing into hash buffer")
			}
		}
	}
}
