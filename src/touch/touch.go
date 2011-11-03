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
        fmt.Fprintln(os.Stderr, "touch: missing file operand")
        os.Exit(1)
    }
    for i := 0; i < flag.NArg(); i++ {
		er := Touch(flag.Arg(i))
		if er != nil {
			fmt.Fprintln(os.Stderr, "touch: cannot create file %s", flag.Arg(i))
		}
    }
}

func Touch(path string) os.Error {
	now := time.Nanoseconds()
	er := os.Chtimes(path, now, now)
	if er != nil {
		_, ew := os.Create(path)
		return ew
	}
	return nil
}
