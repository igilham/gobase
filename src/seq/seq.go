package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
)

const (
	// standard newline character
	Newline = '\n'

	// standard newline character string
	StrNewline = string(Newline)
)

// Seq returns a slice containing a sequence of numbers from start to end in
// step increments. If the following conditions do not hold, Seq returns the
// empty slice.
// step != 0
// if step > 0; end > start
// if step < 0; end < start
func Seq(start, step, end float64, format string) []float64 {
	// are we counting up or down?
	var out []float64
	var direction float64
	switch {
	case step > 0:
		direction = 1.0
	case step < 0:
		direction = -1.0
	default:
		return out
	}
	if (start * direction) > (end * direction) {
		return out
	}
	for num := start; (num * direction) <= (end * direction); num += step {
		out = append(out, num)
	}
	return out
}


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
	var format = flag.String("f", "%g", "")
	start := parseFloat(starts)
	step := parseFloat(steps)
	end := parseFloat(ends)
	out := Seq(start, step, end, *format)
	for n := 0; n < len(out); n++ {
		if out[n] != start {
			fmt.Printf(StrNewline)
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
