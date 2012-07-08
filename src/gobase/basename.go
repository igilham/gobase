package gobase

import (
	"os"
	"strings"
)

// Basename strips the non-directory suffix from a filename.
func Basename(s string) string {
	arr := strings.Split(s, string(os.PathSeparator))
	for i := len(arr) - 1; i >= 0; i-- {
		if len(arr[i]) != 0 {
			s = arr[i]
			break
		}
		if i == 0 {
			return string(os.PathSeparator)
		}
	}
	return s
}
