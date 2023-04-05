package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("io_test/input.txt")
	if err != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got access to it?\n")
		return // exit the function on error
	}
	fileReader := bufio.NewReader(file)
	for true {
		readString, err := fileReader.ReadString('\n')
		println(readString)
		if err == io.EOF {
			return
		}
	}
}
