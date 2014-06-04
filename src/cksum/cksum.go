package main

import (
	"errors"
	"flag"
	"fmt"
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

// prints the checksum  and the size of the checksum in bytes of each file argument
func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		h, e := Cksum(os.Stdin)
		printCksum(h, e)
	} else {
		suppressName := true
		if flag.NArg() > 1 {
			suppressName = false
		}
		for i := 0; i < flag.NArg(); i++ {
			fd, err := os.Open(flag.Arg(i))
			defer fd.Close()
			if err != nil {
				fmt.Fprintln(os.Stderr, "cksum: error reading file ", flag.Arg(i))
			} else {
				h, e := Cksum(fd)
				printNamedCksum(h, fd, e, suppressName)
			}
		}
	}
}

func printCksum(h hash.Hash32, e error) {
	if e != nil {
		fmt.Fprintln(os.Stderr, "cksum: ", e)
	} else if h != nil {
		fmt.Println(fmt.Sprintf("%d %d", h.Sum32(), h.Size()))
	} else {
		fmt.Fprintln(os.Stderr, "cksum: error")
	}
}

func printNamedCksum(h hash.Hash32, fd *os.File, e error, suppressName bool) {
	if e != nil {
		fmt.Fprintf(os.Stderr, "cksum: %s\n", e)
	} else if h != nil {
		if suppressName {
			fmt.Println(h.Sum32())
		} else {
			fmt.Println(fmt.Sprintf("%d %s", h.Sum32(), fd.Name()))
		}
	} else {
		fmt.Fprintln(os.Stderr, "cksum: error")
	}
}
