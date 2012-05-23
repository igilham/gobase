package gobase

import (
	"strings"
)

// Dirname strips the filename from a directory path.
func Dirname(s string) string {
	if s == sep {
		return s
	}
	s = strings.TrimRight(s, sep)
	index := strings.LastIndex(s, sep)
	if index == -1 {
		return cwd
	}
	// remove anything after trailing /
	// cast to []rune avoids multibyte unicode issues
	s = string([]rune(s)[0 : index+1])
	if len(s) <= 1 {
		return s
	}
	return strings.TrimRight(s, sep)
}
