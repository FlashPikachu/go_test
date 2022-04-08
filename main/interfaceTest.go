package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Dog interface {
	eat()
}

func doggo(name *string) {
	fmt.Printf("my name is %s\n", *name)
}

//func main() {
//	//name := "chako"
//	//doggo(&name)
//	//conditionOfDay(7)
//	//day := 100
//	//conditionOfDay(day)
//	//loopTest()
//	//testArrayLoop()
//	//testIota()
//	//forTest()
//	//stringTest()
//	//forRangeTest()
//	//sliceTest()
//	//primaryFunc()
//	//testIn()
//	//testInherit()
//	testChannel()
//}

var values = make(chan int)

func send() {
	rand.Seed(time.Now().UnixNano())
	value := rand.Int31n(10)
	fmt.Printf("send data %d \n", value)
	values <- int(value)
}

func receive() {
	value := <-values
	fmt.Printf("I have reveived your data: %d\n", value)
}

func testChannel() {
	defer close(values)
	go send()

	go receive()
	fmt.Printf("wait...")
	value := <-values
	fmt.Printf("receive : %d\n", value)
	println("end...")
}

type organism struct {
	type_ string
}

type Animal struct {
	organism organism
	name     string
	age      int
}

func (animal Animal) eat(food string) {
	fmt.Printf("eat %s...", food)
}

func (animal Animal) sleep() {
	println("sleep...")
}

func (animal Animal) getInfo() {
	fmt.Printf("my age is %d\nmy name is %s\n", animal.age, animal.name)
}

type Cat struct {
	Animal
	tail bool
}

func testInherit() {
	cat := Cat{Animal: Animal{organism: organism{type_: "female"}, name: "栗山樱原", age: 16}, tail: true}
	cat.getInfo()
	println(cat.name)
	println(cat.Animal.organism.type_)
}

type outputDevice interface {
	write()
	read()
}

func testIn() {
	c := Computer{name: "ibm"}
	c.write()

	var ot outputDevice
	ot = Computer{
		name: "lenevo",
	}
	ot.read()
}

type Computer struct {
	name string
}

func (c Computer) read() {
	//TODO implement me
	println("read")
}

func (c Computer) write() {
	println(c.name)
}

type operator func(int, int) int

func add(int2 int, int3 int) int {
	return int2 + int3
}

func sub(int2 int, int3 int) int {
	return int2 - int3
}

func cal(string2 string) operator {
	switch string2 {
	case "add":
		return add
	case "sub":
		return sub
	default:
		return nil
	}
}

func primaryFunc() {
	ff := cal("add")
	println(ff(3, 1))
	ff = cal("sub")
	println(ff(3, 1))
}

func conditionOfDay(day int) {
	switch day {
	case 1, 2, 3, 4, 5:
		fmt.Println("today is work day!")
	case 6, 7:
		fmt.Println("today is weekend!")
	default:
		fmt.Println("invalid input!")
	}

}

func loopTest() {
	for i := 0; i < 10; i++ {
		println(i)
	}

	r := 1
	for ; r < 10; r++ {
		print(r)
	}

	x := 1
	for x < 10 {
		x++
		println(x)

	}
}

func testArrayLoop() {
	a := [...]int{1, 2, 3, 4}

	for _, i := range a {
		println(i)
	}
}

func testIota() {
	const (
		a1 = iota
		//a2 = iota
		_
		a3 = iota
	)

	const (
		b0 = 20
		b1 = iota
		b2 = iota
	)

	const c = iota

	println(a1)
	//println(a2)
	println(a3)
	println(b0)
	println(b1)
	println(b2)
	println(c)
}

func forTest() {
	for i := 0; true; i++ {
		println(i)
	}
}

func stringTest() {
	s1 := "test1"
	var s2 string = "3"
	println(s2)
	println(s1)

	s3 := `
		<html>text<\html>
		<br>
	`

	println(s3)
}

func forRangeTest() {
	var a = [5]int{1, 2, 3, 4, 5}
	b := [3]string{"a", "b", "c"}
	for index, value := range a {
		fmt.Printf("下标：%d 值：%d\n", index, value)
	}

	for _, value := range b {
		println(value)
	}

	var c = [3]int{0, 1}
	println(c[1])
	fmt.Printf("c: %v", c)
}

func sliceTest() {
	var a = []int{0, 1, 2}
	fmt.Printf("%v", len(a))
	fmt.Printf("%v", cap(a))
	a = append(a, 2, 5, 1, 3, 5)
	for index, val := range a[:] {
		fmt.Printf("\nindex:%d value:%d", index, val)
	}
}
