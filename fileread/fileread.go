/*
A text file reader.
*/
package main

//
// A text file reader.
//
import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fmt.Println("Start....")

	// Open
	// f is *File.
	f, err := os.OpenFile("./data.txt", os.O_RDONLY, 0644)
	if err != nil {
		fmt.Printf("\nOpen Error => %s\n\n", err)
		os.Exit(1)
	}

	// Read lines
	reader := bufio.NewReader(f) // Buffered reader
	for true {
		line, errr := reader.ReadString('\n')
		if errr == io.EOF {
			break
		}
		line = strings.Replace(line, "\n", "", -1) // chomp, sort of ....
		if line == "" {
			break
		}
		fmt.Printf("Next Line Read: |%s|\n", line)
	}

	// Close
	errc := f.Close()
	if errc != nil {
		fmt.Printf("\nClose Error => %s\n\n", errc)
		os.Exit(1)
	}
	//
	fmt.Println("End....")
}
