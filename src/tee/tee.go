package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

const (
	ARBUF = 512 // Arrays buffer size
	CHBUF = 1   // Channels buffer size
)

//var buffer []byte
var dontCreate = flag.Bool("a", false, "Append the output to the files, instead of overwrite it")

func main() {
	flag.Parse()
	Channels := make([]chan []byte, flag.NArg()+1) // Makes the channel array
	Files := make([]*os.File, flag.NArg()+1)       // Makes the file array
	defer cleanupFiles(Files)

	// The first output file is stdout
	Files[0] = os.Stdout
	Channels[0] = make(chan []byte)
	go writeToFile(Files[0], Channels[0])

	for i := 1; i < len(Files); i++ {
		var err error
		if *dontCreate {
			Files[i], err = os.OpenFile(flag.Arg(i-1), os.O_APPEND, 0664)
			if err != nil {
				if err == os.ErrNotExist {
					var nerr error
					Files[i], nerr = os.Create(flag.Arg(i - 1))
					if nerr != nil {
						log.Fatal(nerr)
					}
				} else {
					log.Fatal(err)
				}
			}
		} else {
			Files[i], err = os.Create(flag.Arg(i - 1))
		}
		if err != nil {
			fmt.Fprintf(os.Stderr,
				"tee: cannot create file - %s - %s\n",
				flag.Arg(i-1),
				err.Error())
		}
		Channels[i] = make(chan []byte, CHBUF) // Makes the channel
		go writeToFile(Files[i], Channels[i])  // Starts the writers
	}
SL:
	for buffer := make([]byte, ARBUF); ; { // SL stands for Send Loop
		end, ok := os.Stdin.Read(buffer)
		for i := 0; i < len(Channels); i++ {
			Channels[i] <- buffer[:end]
		}
		if ok == io.EOF {
			for i := 0; i < len(Channels); i++ {
				close(Channels[i])
			}
			break SL
		}

	}
}

// Writes to the file argument
func writeToFile(file *os.File, channel chan []byte) {
	defer file.Close() // Closes the file
	for {
		select {
		case buf, open := <-channel:
			if !open {
				return
			}
			file.Write(buf)
		}
	}
}

func cleanupFiles(files []*os.File) {
	for _, f := range files {
		f.Close()
	}
}
