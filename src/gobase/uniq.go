package gobase

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)


// Uniq prints the input, omitting duplicated lines
func Uniq(fd *os.File) error {
	var prevLine string
	reader := bufio.NewReader(fd)
	buffer := new(bytes.Buffer)
	firstLine := true
	for {
		part, prefix, er := reader.ReadLine()
		if er != nil {
			break
		}
		buffer.Write(part)
		if !prefix {
			thisLine := buffer.String()
			if firstLine || prevLine != thisLine {
				fmt.Println(thisLine)
				firstLine = false
			}
			prevLine = thisLine
			buffer.Reset()
		}
	}
	return nil
}