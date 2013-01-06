package gobase

import (
	"bufio"
	"bytes"
	"os"
)

// Head gets the first n lines of a file. If n == 0, then attempt to get all 
// the lines.
func Head(file *os.File, n int) []string {
	var lines []string
	var part []byte
	var prefix bool
	var err error
	reader := bufio.NewReader(file)
	buffer := new(bytes.Buffer)
	for n != len(lines) || n == 0 {
		if part, prefix, err = reader.ReadLine(); err != nil {
			break
		}
		buffer.Write(part)
		if !prefix {
			lines = append(lines, buffer.String())
			buffer.Reset()
		}
	}
	return lines
}
