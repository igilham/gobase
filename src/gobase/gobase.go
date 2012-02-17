package gobase

import (
	"os"
	"strings"
	"time"
)

const (
	cwd string = "."
	sep = string(os.PathSeparator)
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

func Touch(path string) os.Error {
	now := time.Nanoseconds()
	er := os.Chtimes(path, now, now)
	if er != nil {
		_, ew := os.Create(path)
		return ew
	}
	return nil
}
