package gobase

import (
	"bytes"
	"bufio"
	"hash"
	"hash/crc32"
	"os"
	"strings"
	"time"
	"utf8"
)

const (
	cwd string = "."
	sep = string(os.PathSeparator)
	newline = '\n'
	buf_size = 4096
)

var (
	whiteSpaceChars = []byte{9, 10, 11, 12, 13, 32}
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

// Cat concatenate files.
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

// Cksum calculates a simple checksum for a file. Currently uses crc32.
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

// Dirname strips the filename from a directory path.
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

// FileExists returns true if path points to an existing file.
func FileExists(path string) bool {
	fi, _ := os.Stat(path)
	return fi != nil
}
// Head gets the first n lines of a file. If n == 0, then attempt to get all 
// the lines.
func Head(file *os.File, n int) ([]string) {
	var lines []string
	var part []byte
	var	prefix bool
	var err os.Error
	reader := bufio.NewReader(file)
	buffer := bytes.NewBuffer(make([]byte, buf_size))
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

// Touch updates the access and modify times of a file, or creates a new
// file if it doesn't already exist.
func Touch(path string) os.Error {
	now := time.Nanoseconds()
	er := os.Chtimes(path, now, now)
	if er != nil {
		_, ew := os.Create(path)
		return ew
	}
	return nil
}

// WordCount contains the counts calculated by the Wc algorithm.
type WordCount struct {
	Bytes uint64
	Chars uint64
	Words uint64
	Lines uint64
}

// utf8Point determines if a byte is a valid UTF8 code point. Copied
// from sbase/wc.c (suckless.org).
func utf8Point(c byte) bool {
	return c & 0xc0 != utf8.RuneSelf
}

// isSpace determines if a byte contains a space character.
func isSpace(c byte) bool {
	for _, b := range whiteSpaceChars {
		if c == b {
			return true
		}
	}
	return false
}

// Wc counts the lines, words, bytes and characters in a file.
func Wc(fd *os.File) (WordCount, os.Error) {
	reader := bufio.NewReader(fd)
	wc := WordCount{0, 0, 0, 0}
	word := false
	var c byte
	var e os.Error
	for c, e = reader.ReadByte(); e == nil; c, e = reader.ReadByte() {
		wc.Bytes++
		if utf8Point(c) {
			wc.Chars++
		}
		if c == newline {
			wc.Lines++
		}
		if !isSpace(c) {
			word = true
		} else if word {
			word = false
			wc.Words++
		}
	}
	if e == os.EOF {
		return wc, nil
	}
	return wc, e
}
