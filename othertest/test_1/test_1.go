package main

import "fmt"

type User struct {
	name string
	age  int
}

func (u User) sayName() {
	fmt.Printf("my name is %v\n", u.name)
}

func main() {
	if a := 1; false {
	} else if b := 2; false {
	} else if c := 3; false {
	} else {
		println(a, b, c)
	}

	user1 := User{
		name: "pharsalia",
		age:  10,
	}

	user1.sayName()
}
