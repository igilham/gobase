package main

import (
    "flag"
    "fmt"
    "os"
    "os/user"
)

// Prints the effective user ID.
func main() {
    flag.Parse()
    user, error := user.LookupId(os.Geteuid())
    if (error != nil) {
        fmt.Fprintln(os.Stderr, "whoami: %s", error.String())
        os.Exit(1)
    }
    fmt.Println(user.Username);
}
