package main

import (
    "flag"
    "fmt"
    "os"
)

func cat(fd *os.File) {
    const NBUF = 512
    var buf [NBUF]byte
    for {
        switch nr, er := fd.Read(buf[:]); true {
            case nr < 0:
                fmt.Fprintf(os.Stderr, 
                    "cat: error reading from %s: %s\n",
                    fd, er.String())
                os.Exit(1)
            case nr == 0: // EOF
                return
            case nr > 0:
                if nw, ew := os.Stdout.Write(buf[0:nr]); nw != nr {
                    fmt.Fprintf(os.Stderr, "cat: error writing from %s: %s\n", fd, ew.String())
                    os.Exit(1)
                }
        }
    }
}

// concatenate the specified files, joining with a newline. If -n is specified, 
// the newline will be omitted.
func main() {
    flag.Parse()
    if flag.NArg() == 0 {
        cat(os.Stdin)
    }
    for i := 0; i < flag.NArg(); i++ {
        fd, err := os.Open(flag.Arg(i))
        if fd == nil {
            fmt.Fprintf(os.Stderr, "cat: can't open %s: error %s\n", flag.Arg(i), err)
            os.Exit(1)
        }
        defer fd.Close()
        cat(fd)
    }
}
