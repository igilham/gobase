package gobase

import (
	"os"
	"strings"
)

// Dirname strips the filename from a directory path.
func Dirname(s string) string {
	if s == string(os.PathSeparator) {
		return s
	}
	s = strings.TrimRight(s, string(os.PathSeparator))
	index := strings.LastIndex(s, string(os.PathSeparator))
	if index == -1 {
		return cwd
	}
	// remove anything after trailing /
	// cast to []rune avoids multibyte unicode issues
	s = string([]rune(s)[0 : index+1])
	if len(s) <= 1 {
		return s
	}
	return strings.TrimRight(s, string(os.PathSeparator))
}
