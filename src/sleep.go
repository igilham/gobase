// sleep.go

package main

import (
    "flag"
    "fmt"
    "os"
    "time"
    "strconv"
)

func usage() {
    fmt.Printf("usage: sleep MILLISECONDS\n")
}

// BUG: there is an error in here somewhere. It compiles without warnings
// but does not correctly sleep
func main() {
    flag.Parse()
    if flag.NArg() != 1 {
        fmt.Printf("sleep: invalid arguments\n")
        os.Exit(1)
    }
    str := flag.Arg(0)
    sleeptime, er := strconv.Atoi64(str)
    if er != nil {
        usage()
    } else {
        // BUG: sleep doesn't work
        time.Sleep(sleeptime)
    }
}

