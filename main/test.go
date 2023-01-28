package main

import (
	"fmt"
	"time"
)

//func main() {
//	fmt.Println("hello world")
//}
type user struct {
	id         int
	score      float32
	enrollment time.Time
	name, addr string
}

var u user

func main() {
	u = user{
		id:         1,
		score:      99,
		enrollment: time.Now(),
		name:       "yuu",
		addr:       "hebudilii",
	}
	fmt.Println(u)

	str := `line1 \n line2 \n line3`
	str2 := "line1 \nline2 \nline3"
	println(str)
	println(str2)
	str3 := `test`
	fmt.Printf("%v\n", str3[len(str3)-1])
	var arrAge = [4]int{1, 2, 3, 4}
	fmt.Printf("%v", arrAge)
}
