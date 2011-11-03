package main

import (
	"os"
	"flag"
	"fmt"
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

	// The first output file is stdout
	Files[0] = os.Stdout
	Channels[0] = make(chan []byte)
	go writeToFile(Files[0], Channels[0])

	for i := 1; i < len(Files); i++ {
		var err os.Error
		if *dontCreate {
			Files[i], err = os.OpenFile(flag.Arg(i-1), os.O_APPEND, 0664)
			if err != nil {
				if err == os.ENOENT {
					var nerr os.Error
					Files[i], nerr = os.Create(flag.Arg(i-1))
					if nerr != nil {
						fmt.Fprintln(os.Stderr, nerr)
						os.Exit(1)
					}
				} else {
					fmt.Fprintln(os.Stderr, err)
					os.Exit(1)
				}
			}
		} else {
			Files[i], err = os.Create(flag.Arg(i-1))
		}
		if err != nil {
			fmt.Fprintln(os.Stderr,
				"tee: cannot create file - %s - %s",
				flag.Arg(i-1),
				err.String())
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
		if ok == os.EOF {
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

// Splits an input channel to an array of channels
func split(inputchan chan []byte, chanar []chan []byte) { // chanar stands for channel array
	for {
		buffer, open := <-inputchan
		if !open {
			for i := 0; i <= len(chanar); i++ {
				close(chanar[i])
			}
			return
		}
		for i := 0; i <= len(chanar); i++ {
			chanar[i] <- buffer
		}
	}
}
