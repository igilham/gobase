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

// Prints the effective user ID.
func main() {
    flag.Parse()
    user, error := user.LookupId(os.Geteuid())
    if (error != nil) {
        fmt.Printf(error.String() + newline)
        os.Exit(1)
    }
    fmt.Printf(user.Username + newline);
}
