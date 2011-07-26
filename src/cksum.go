// cksum.go

package main

import (
    "flag"
    "fmt"
    "hash/crc64"
    "os"
)

// prints the checksum of each file argument
func main() {
    flag.Parse()
    if flag.NArg() == 0 {
        fmt.Printf("no args\n")
    } else {
        for i := 0; i < flag.NArg(); i++ {
            cksum(flag.Arg(i))
        }
    }
}

func cksum(s string) {
    table := crc64.MakeTable(crc64.ISO)
    h := crc64.New(table)
    
    fd, err := os.Open(s)
    if (err != nil) {
        fmt.Printf("error reading file %s\n", s)
        return
    }
    
    const NBUF = 512
    var buf [NBUF]byte
    for {
        switch nr, er := fd.Read(buf[:]); true {
            case nr < 0:
                fmt.Printf("error reading from %s: %s\n", fd, er.String())
                return
            case nr == 0: // EOF
                fmt.Printf("%d %s\n", h.Sum64(), s)
                return
            case nr > 0:
                _, ew := h.Write(buf[0:nr])
                if ew != nil {
                    fmt.Printf("error writing into hash buffer\n")
                    return
                }
        }
    }
}

