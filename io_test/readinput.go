package main

import "fmt"

var (
	firstName, lastName, s string
	i                      int
	f                      float32
	input                  = "56.12 / 5212 / Go"
	format                 = "%f / %d / %s"
)

func main() {
	var temp, temp1, temp2 string
	fmt.Println("Please enter your words: ")
	fmt.Scanln(&temp, &temp1, &temp2)

	fmt.Printf("temp: %s, temp2: %s, temp3: %s", temp, temp1, temp2)
	fmt.Sscanf(input, format, &f, &i, &s)
	fmt.Println("From the string we read: ", f, i, s)
}
