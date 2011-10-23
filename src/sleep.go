package main

import (
	"flag"
	"fmt"
	"os"
	"time"
	"strconv"
)

var factor float64
var number int64
var amountOfTime int64

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		fail()
	}
	for i := 0; i < flag.NArg(); i++ {
		switch flag.Arg(0)[len(flag.Arg(0))-1] {
		case 'm':
			factor = 6e10
			number, _ = strconv.Atoi64(flag.Arg(i)[:len(flag.Arg(i))-1])
		case 'h':
			factor = 36e11
			number, _ = strconv.Atoi64(flag.Arg(i)[:len(flag.Arg(i))-1])
		case 'd':
			factor = 864e11
			number, _ = strconv.Atoi64(flag.Arg(i)[:len(flag.Arg(i))-1])
		case 's':
			factor = 1e9
			number, _ = strconv.Atoi64(flag.Arg(i)[:len(flag.Arg(i))-1])
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			factor = 1e9
			number, _ = strconv.Atoi64(flag.Arg(i))
		default:
			fail()

		}
		amountOfTime += number * int64(factor)

	}
	time.Sleep(amountOfTime)
}
func fail() {
	fmt.Println("sleep: invalid arguments")
	os.Exit(1)
}
