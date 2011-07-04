// touch.go

package main

import (
    "flag"
    "fmt"
    "os"
    "time"
)

func main() {
    flag.Parse()
    for i := 0; i < flag.NArg(); i++ {
        now := time.Nanoseconds()
        err := os.Chtimes(flag.Arg(i), now, now)
        if err != nil {
            fmt.Printf("touch: cannot acccess file " + flag.Arg(i))
        }
    }
}
