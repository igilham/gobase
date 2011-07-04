// ls.go
// TODO: come back to this and fix it later

package main

import (
    "os"
    "io"
    "flag"
    "fmt"
)

var longList = flag.Bool("l", false, "use a long listing format")
var all = flag.Bool("a", false, "do not ignore entries starting with .")
var almostAll = flag.Bool("A", false, "do not list implied . and ..")
var ignoreBackups = flag.Bool("B", true, "do not list implied entries ending with ~")

func longListheader() string {
    return "NAME"
}

func closeAll(items []io.Closer) {
    for i := 0; i < len(items); i++ {
        items[i].Close()
    }
}

func main() {
    flag.Parse()
    var c = flag.NArg()
    bound := c + 1
    var toScan [bound]os.File
    if c > 0 {
        for i := 0; i < c; i++ {
            thing, err := os.Open(flag.Arg(i))
            toScan[i] = thing
        }
    } else {
        thing, err := os.Open(".")
        toScan[thing]
    }
    defer closeAll(toScan)
    for i := 0; i < len(toScan); i++ {
        if c > 0 {
            fmt.Printf("in" + toScan[i].Name)
        }
        var items []string = toScan[i].Readdirnames(1 << 31)
        defer closeAll(items)
        if *longList {
            fmt.Printf(longListHeader())
        }
        for j := 0; j < len(items); j++ {
            if *longList {
                fmt.Printf(items[j].Name + "\n")
            } else {
                fmt.Printf(items[j].Name + "\t")
            }
        }
    }
}

