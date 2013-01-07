package main

import (
	"flag"
	"log"
	"time"
)

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		fail()
	}
	d, e := time.ParseDuration(flag.Arg(0))
	if e != nil {
		fail()
	}
	time.Sleep(d)
}

func fail() {
	log.Fatal("sleep: invalid arguments")
}
