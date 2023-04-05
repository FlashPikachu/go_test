package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)

	fmt.Println("Please input your contents:")

	readString, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("read error")
	}
	fmt.Printf("your input contents are: %s", readString)
}
