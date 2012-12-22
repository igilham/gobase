package gobase

import (
	"fmt"
	"os"
)

func Pwd() {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, "pwd: ", err)
		os.Exit(1)
	} else {
		fmt.Println(wd)
	}
}