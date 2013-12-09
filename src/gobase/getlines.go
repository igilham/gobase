package gobase

import (
	"bufio"
	"bytes"
	"os"
)

// GetLines reads the first n lines of the file f.
// The end result is similar to Head, but suited to different use cases.
// It returns a slice of strings for each line read and an error if encountered.
func GetLines(file *os.File, n int) ([]string, error) {
	var lines []string
	var part []byte
	var prefix bool
	var err error
	reader := bufio.NewReader(file)
	buffer := new(bytes.Buffer)
	for n != len(lines) || n == 0 {
		if part, prefix, err = reader.ReadLine(); err != nil {
			return lines, err
		}
		buffer.Write(part)
		if !prefix {
			lines = append(lines, buffer.String())
			buffer.Reset()
		}
	}
	return lines, nil
}
