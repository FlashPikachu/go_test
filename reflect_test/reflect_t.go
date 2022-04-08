package main

import (
	"fmt"
	"reflect"
)

func reflectType(x interface{}) {
	type_ := reflect.TypeOf(x)
	fmt.Printf("Type:%v Name:%v kind:%v\n", type_, type_.Name(), type_.Kind())

}

type Animal interface {
	bar()
}

type cat struct {
	name string
}

func (c cat) bar() {
	fmt.Printf("my name is %s\n", c.name)
}

func main() {
	cat := cat{name: "绯鞠"}
	cat.bar()
	reflectType(cat)
	var t1 []int
	reflectType(t1)
}
