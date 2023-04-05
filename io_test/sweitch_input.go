package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your name:")
	readString, err := inputReader.ReadString('\n')
	if err == nil {
		println("your name is : " + readString)
		switch readString {
		case "pharsalia\r\n":
			fmt.Println("Welcome ! " + readString)
		case "tthr\r\n":
			fmt.Println("Welcome " + readString)
		default:
			fmt.Println("Who are you?")
		}
	}
}
