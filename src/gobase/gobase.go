package gobase

import (
	"os"
)

const (
	cwd string      = "."
	sep             = string(os.PathSeparator)
	newline         = '\n'
	strNewline		= string(newline)
	buf_size        = 4096
)

var (
	whiteSpaceChars = []byte{9, 10, 11, 12, 13, 32}
)

// FileExists returns true if path points to an existing file.
func FileExists(path string) bool {
	fi, _ := os.Stat(path)
	return fi != nil
}
