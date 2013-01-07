package gobase

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

// Head gets the first n lines of a file. If n == 0, then attempt to get all 
// the lines.
func Head(file *os.File, n int) error {
	lines := 0
	var part []byte
	var prefix bool
	var err error
	reader := bufio.NewReader(file)
	buffer := new(bytes.Buffer)
	for n != lines || n == 0 {
		if part, prefix, err = reader.ReadLine(); err != nil {
			return err
		}
		buffer.Write(part)
		if !prefix {
			fmt.Printf("%s\n", buffer)
			buffer.Reset()
			lines++
		}
	}
	return nil
}
