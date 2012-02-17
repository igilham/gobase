package gobase

import (
	"bytes"
	"bufio"
	"hash"
	"hash/crc32"
	"os"
	"strings"
	"time"
)

const (
	cwd string = "."
	sep = string(os.PathSeparator)
	buf_size = 4096
)

// strip the non-directory suffix from a filename
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

// concatenate files
func Cat(fd *os.File) os.Error {
    const NBUF = 512
    var buf [NBUF]byte
    for {
        switch nr, er := fd.Read(buf[:]); true {
            case nr < 0:
                return os.NewError("error reading from " + fd.Name() + " : " + er.String())
            case nr == 0: // EOF
                return nil
            case nr > 0:
                if nw, _ := os.Stdout.Write(buf[0:nr]); nw != nr {
                    return os.NewError("error reading from " + fd.Name() + " : " + er.String())
                }
        }
    }
    return nil
}

// checksum a file
func Cksum(fd *os.File) (hash.Hash32, os.Error) {
    table := crc32.MakeTable(crc32.Castagnoli)
    h := crc32.New(table)
    const NBUF = 512
    var buf [NBUF]byte
    for {
        switch nr, er := fd.Read(buf[:]); true {
        case nr < 0:
            return nil, os.NewError("error reading from " + fd.Name() + " : " + er.String())
        case nr == 0: // EOF
            break
        case nr > 0:
            _, ew := h.Write(buf[0:nr])
            if ew != nil {
                return nil, os.NewError("error writing into hash buffer")
            }
        }
    }
    return h, nil
}

// strip the filename from a directory
func Dirname (s string) string {
    if s == sep {
        return s
    }
    s = strings.TrimRight(s, sep)
    index := strings.LastIndex(s, sep)
    if index == -1 {
        return cwd
    }
    // remove anything after trailing /
    // cast to []int avoids multibyte unicode issues
    s = string([]int(s)[0:index+1])
    if len(s) <= 1 {
		return s
    }
    return strings.TrimRight(s, sep)
}

func FileExists(path string) bool {
	fi, _ := os.Stat(path)
	return fi != nil
}

func Head(file *os.File, n int) ([]string) {
	var lines []string
	var part []byte
	var	prefix bool
	var err os.Error
	reader := bufio.NewReader(file)
	buffer := bytes.NewBuffer(make([]byte, buf_size))
	for n != len(lines) {
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

func Touch(path string) os.Error {
	now := time.Nanoseconds()
	er := os.Chtimes(path, now, now)
	if er != nil {
		_, ew := os.Create(path)
		return ew
	}
	return nil
}
