package main

import (
	"../gobase"
	"flag"
	"fmt"
	"log"
	"strconv"
)

var format = flag.String("f", "%g", "")

// seq prints a sequence of numbers
func main() {
	flag.Parse()
	argInd := 0
	starts := "1.0"
	steps := "1.0"
	var ends string
	switch flag.NArg() {
	case 3:
		starts = flag.Arg(argInd)
		argInd++
		steps = flag.Arg(argInd)
		argInd++
		ends = flag.Arg(argInd)
	case 2:
		starts = flag.Arg(argInd)
		argInd++
		fallthrough
	case 1:
		ends = flag.Arg(argInd)
	default:
		usageExit()
	}
	start := parseFloat(starts)
	step := parseFloat(steps)
	end := parseFloat(ends)
	out := gobase.Seq(start, step, end)
	for n := 0; n < len(out); n++ {
		if out[n] != start {
			fmt.Printf(gobase.StrNewline)
		}
		fmt.Printf(*format, out[n])
	}
}

func parseFloat(s string) float64 {
	const bitSize = 64
	f, e := strconv.ParseFloat(s, bitSize)
	if e != nil && e.Error() == "ErrSyntax" {
		usageExit()
	}
	return f
}

func usageExit() {
	log.Fatalf("usage:\nseq [-f FORMAT] end\nseq [-f FORMAT] start end\nseq [-f FORMAT] start step end")
}
