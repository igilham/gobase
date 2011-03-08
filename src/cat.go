// cat.go

package coreutils.cat

import (
	"os"
	"flag"
	"bufio"
)

var omitNewline = flag.Bool("n", false, "don't print final newline")
const Newline = "\n"

func readFile(string fname) string {
	file, err := os.Open(fname, os.O_RDONLY, 0666)
	defer file.Close()
	if err != nill {
		panic(err)
	}
	var s string = ""
	read := bufio.NewReader(file)
	for {
		str, err := read.ReadString('\n')
		if err != nill {
			if err != os.EOF {
				panic(err)
			}
			break
		}
		s += str
	}
	return s
}

/* concatenate the specified files, joining with a newline. If -n is specified, 
the newline will be omitted. 
func main() {
	flag.Parse()
	var s string = ""
	if flag.NArg() == 0 {
		s += readFile(os.Stdin)
	}
	else {
		for i := 0; i < flag.NArg(); i++ {
			s += readFile(flag.Arg(i))
			if !omitNewline {
				s += Newline
			}
		}
	}
	os.Stdout.WriteString(s)
}