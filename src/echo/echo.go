package main

import (
	"flag"
	"fmt"
    "strings"
)

var omitNewline = flag.Bool("n", false, "don't print final newline")

const (
    Space = " "
    Newline = "\n"
)

/* Echo the arguments to Stdout. If -n is specified, omit a trailing newline.*/
func main() {
    flag.Parse()
    s := strings.Join(flag.Args(), Space)
    if !*omitNewline {
        s += Newline
    }
    fmt.Printf(s)
}
