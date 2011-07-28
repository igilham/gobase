// touch.go

package main

import (
    "flag"
    "fmt"
    "os"
    "time"
)

// Change file timestamps or create files that don't already exist
func main() {
    flag.Parse()
    if flag.NArg() == 0 {
        fmt.Printf("touch: missing file operand\n")
        os.Exit(1)
    }
    for i := 0; i < flag.NArg(); i++ {
        now := time.Nanoseconds()
        err := os.Chtimes(flag.Arg(i), now, now)
        if err != nil {
            _, ew := os.Create(flag.Arg(i))
            if ew != nil {
                fmt.Printf("touch: cannot create file %s\n", flag.Arg(i))
            }
        }
    }
}

