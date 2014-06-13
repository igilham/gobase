package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"unicode/utf8"
)

// WordCount contains the counts calculated by the Wc algorithm.
type WordCount struct {
	Bytes uint64
	Chars uint64
	Words uint64
	Lines uint64
}

const (
	// standard newline character
	Newline = '\n'
)

var (
	countLines = flag.Bool("l", false, "count lines")
	countWords = flag.Bool("w", false, "count words")
	countBytes = flag.Bool("c", false, "count bytes")
	countChars = flag.Bool("m", false, "count characters")

	// slice of characters representative of whitespace
	whiteSpaceChars = []byte{9, 10, 11, 12, 13, 32}
)

// utf8Point determines if a byte is a valid UTF8 code point. Copied
// from sbase/wc.c (suckless.org).
func utf8Point(c byte) bool {
	return c&0xc0 != utf8.RuneSelf
}

// isSpace determines if a byte contains a space character.
func isSpace(c byte) bool {
	for _, b := range whiteSpaceChars {
		if c == b {
			return true
		}
	}
	return false
}

// Wc counts the lines, words, bytes and characters in a file.
func Wc(fd *os.File) (WordCount, error) {
	reader := bufio.NewReader(fd)
	wc := WordCount{0, 0, 0, 0}
	word := false
	var c byte
	var e error
	for c, e = reader.ReadByte(); e == nil; c, e = reader.ReadByte() {
		wc.Bytes++
		if utf8Point(c) {
			wc.Chars++
		}
		if c == Newline {
			wc.Lines++
		}
		if !isSpace(c) {
			word = true
		} else if word {
			word = false
			wc.Words++
		}
	}
	if e == io.EOF {
		return wc, nil
	}
	return wc, e
}

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		wc(os.Stdin)
	}
	for i := 0; i < flag.NArg(); i++ {
		fd, er := os.Open(flag.Arg(i))
		if er != nil {
			fmt.Fprintf(os.Stderr, "wc: %v\n", er)
			os.Exit(1)
		}
		defer fd.Close()
		wc(fd)
	}
}

func wc(fd *os.File) {
	wc, ew := Wc(fd)
	if ew != nil {
		fmt.Fprintf(os.Stderr, "wc: %v\n", ew)
		os.Exit(1)
	} else {
		output(wc, fd.Name())
	}
}

func output(wc WordCount, name string) {
	noFlags := !*countLines && !*countWords && !*countBytes && !*countChars
	oneFlag := !(*countLines && *countWords) &&
		!(*countLines && *countBytes) &&
		!(*countLines && *countChars) &&
		!(*countWords && *countBytes) &&
		!(*countWords && *countChars) &&
		(*countLines || *countWords || *countBytes || *countChars)
	format := " %5d"
	if oneFlag {
		format = "%d"
	}
	if *countLines || noFlags {
		fmt.Printf(format, wc.Lines)
	}
	if *countWords || noFlags {
		fmt.Printf(format, wc.Words)
	}
	if *countBytes || noFlags {
		fmt.Printf(format, wc.Bytes)
	}
	if *countChars {
		fmt.Printf(format, wc.Chars)
	}
	fmt.Printf(" %s\n", name)
}
