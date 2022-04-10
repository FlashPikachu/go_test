package main

import (
	"fmt"
	"reflect"
)

func reflectType(x interface{}) {
	type_ := reflect.TypeOf(x)
	fmt.Printf("Type:%v Name:%v kind:%v\n", type_, type_.Name(), type_.Kind())

}

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	fmt.Printf("%v,%T\n", v, v)
}

func reflectSetValue(x interface{}) {
	v := reflect.ValueOf(x)
	kind := v.Elem().Kind()
	fmt.Printf("kind : %v\n", kind)
	//vAddr := &v
	switch kind {
	case reflect.Int32:
		v.Elem().SetInt(100)
	}

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
	//cat := cat{name: "绯鞠"}
	//cat.bar()
	//reflectType(cat)
	//var t1 []int
	//reflectType(t1)

	//set value
	var t int32 = 20
	reflectSetValue(&t)
	println(t)
}
