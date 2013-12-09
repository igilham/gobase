package gobase

import (
	"os"
	"time"
)

// Touch updates the access and modify times of a file, or creates a new
// file if it doesn't already exist.
func Touch(path string) error {
	now := time.Now()
	er := os.Chtimes(path, now, now)
	if er != nil {
		_, ew := os.Create(path)
		return ew
	}
	return nil
}
