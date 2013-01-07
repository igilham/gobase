package gobase

import (
	"fmt"
	"log"
	"os"
)

func Pwd() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("pwd: %v", err)
	} else {
		fmt.Println(wd)
	}
}
