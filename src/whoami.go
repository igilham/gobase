// whoami.go

package main

import (
    "flag"
    "fmt"
    "os"
    "os/user"
)

const (
    newline = "\n"
)

func main() {
    flag.Parse()
    user, error := user.LookupId(os.Geteuid())
    if (error == nil) {
        fmt.Printf(user.Username + newline);
    } else {
        fmt.Printf(error.String() + newline)
    }
}
