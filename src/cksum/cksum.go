package main

import (
    "flag"
    "fmt"
    "hash"
    "os"
    "gobase"
)

// prints the checksum  and the size of the checksum in bytes of each file argument
func main() {
    flag.Parse()
    if flag.NArg() == 0 {
		h, e := gobase.Cksum(os.Stdin)
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
				h, e := gobase.Cksum(fd)
                printNamedCksum(h, fd, e, suppressName)
            }
        }
    }
}

func printCksum(h hash.Hash32, e os.Error) {
	if e != nil {
		fmt.Fprintln(os.Stderr, "cksum: ", e)
	} else if h != nil {
		fmt.Println(fmt.Sprintf("%d %d", h.Sum32(), h.Size()))
	} else {
		fmt.Fprintln(os.Stderr, "cksum: error")
	}
}

func printNamedCksum(h hash.Hash32, fd *os.File, e os.Error, suppressName bool) {
	if e != nil {
		fmt.Fprintln(os.Stderr, "cksum: %s", e)
	} else if h != nil {
		if (suppressName) {
			fmt.Println(h.Sum32())
		} else {
			fmt.Println(fmt.Sprintf("%d %s", h.Sum32(), fd.Name()))
		}
	} else {
		fmt.Fprintln(os.Stderr, "cksum: error")
	}
}
