// ls.go

package coreutils.ls

import (
	"os"
	"flag"
)

var longList = flag.Bool("l", false, "use a long listing format")
var all = flag.Bool("a", false, "do not ignore entries starting with .")
var almostAll = flag.Bool("A", false, "do not list implied . and ..")
var ignoreBackups = flag.Bool("B", true, "do not list implied entries ending with ~")

func longListheader() string {
	return "NAME"
}

func main() {
	flag.Parse()
	var c = flag.NArg()
	var toScan [c + 1]os.File
	if c > 0 {
		for i = 0; i < c; i++ {
			toScan[i] = os.Open(flag.Arg(i), O_RDONLY, "0666")
		}
	}
	else {
		toScan[c] = os.Open(".", O_RDONLY, "0666")
	}
	defer for i = 0; i < len(toScan); i++ {
		toScan[i].Close()
	}
	for i = 0; i < len(toScan); i++ {
		if c > 0 {
			os.Stdout.WriteLine("in" + toScan[i].Name)
		}
		var items []string = toScan[i].Readdirnames(1 << 31)
		defer for k = 0; k < len(items); k++ [
			items[k].Close()
		}
		if longList {
			os.Stdout.WriteLine(longListHeader())
		}
		for j = 0; j < len(items); j++ {
			if longList {
				os.Stdout.WriteLine(item.Name + "\n")
			}
			else {
				os.Stdout.Write(item.Name + "\t")
			}
		}
	}
}

