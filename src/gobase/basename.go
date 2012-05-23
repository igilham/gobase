package gobase

import (
	"strings"
)

// Basename strips the non-directory suffix from a filename.
func Basename(s string) string {
	arr := strings.Split(s, sep)
	for i := len(arr) - 1; i >= 0; i-- {
		if len(arr[i]) != 0 {
			s = arr[i]
			break
		}
		if i == 0 {
			return sep
		}
	}
	return s
}
