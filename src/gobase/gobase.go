package gobase

// This file contains utility constants, variables and functions used in the
// gobase package.

import (
	//"bufio"
	//"bytes"
	"os"
)

const (
	// string representing the current path in the operating system
	cwd string      = "."
	
	// file path separator
	sep = string(os.PathSeparator)
	
	// standard newline character
	Newline         = '\n'
	
	// standard newline character string
	StrNewline		= string(Newline)
	
	// common buffer size for file reading functions
	buf_size        = 4096
)

var (
	// slice of characters representative of whitespace
	whiteSpaceChars = []byte{9, 10, 11, 12, 13, 32}
)

// FileExists returns true if path points to an existing file, and false
// otherwise.
func FileExists(path string) bool {
	fi, _ := os.Stat(path)
	return fi != nil
}
