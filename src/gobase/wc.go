package gobase

import (
	"bufio"
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
