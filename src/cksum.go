// cksum.go

package main

import (
    "flag"
    "fmt"
    "hash"
    "hash/crc32"
    "os"
)

// prints the checksum  and the size of the checksum in bytes of each file argument
func main() {
    flag.Parse()
    if flag.NArg() == 0 {
        if h := cksum(os.Stdin); h != nil {
            fmt.Printf("%d %d\n", h.Sum32(), h.Size())
        }
    } else {
        supressName := true
        if flag.NArg() > 1 {
            supressName = false
        }
        for i := 0; i < flag.NArg(); i++ {
            if fd, err := os.Open(flag.Arg(i)); err != nil {
                fmt.Printf("error reading file %s\n", flag.Arg(i))
            } else {
                if h := cksum(fd); h != nil {
                    if (supressName) {
                        fmt.Printf("%d\n", h.Sum32())
                    } else {
                        fmt.Printf("%d %s\n", h.Sum32(), fd.Name())
                    }
                }
            }
        }
    }
}

func cksum(fd *os.File) hash.Hash32 {
    table := crc32.MakeTable(crc32.Castagnoli)
    h := crc32.New(table)
    const NBUF = 512
    var buf [NBUF]byte
    for {
        switch nr, er := fd.Read(buf[:]); true {
        case nr < 0:
            fmt.Printf("error reading from %s: %s\n", fd, er.String())
            return nil
        case nr == 0: // EOF
            //fmt.Printf("%d %d %s\n", h.Sum32(), h.Size(), fd.Name())
            return h
        case nr > 0:
            _, ew := h.Write(buf[0:nr])
            if ew != nil {
                fmt.Printf("error writing into hash buffer\n")
                return nil
            }
        }
    }
    return nil
}

