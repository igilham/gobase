package gobase

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

func Uniq(fd *os.File) error {
	var prevLine []byte
	reader := bufio.NewReader(fd)
	buffer := new(bytes.Buffer)
	for {
		part, prefix, er := reader.ReadLine()
		if er == io.EOF {
			break
		} else if er != nil {
			return er
		}
		buffer.Write(part)
		if !prefix {
			thisLine := buffer.Bytes()
			if bytes.Compare(thisLine, prevLine) != 0 {
				if len(prevLine) != 0 {
					if _, ew := fmt.Fprintf(os.Stdout, strNewline); ew != nil {
						return ew
					}
				}
				if _, ew := os.Stdout.Write(thisLine); ew != nil {
					return ew
				}
				prevLine = thisLine
			}
			buffer.Reset()
		}
	}
	return nil
}