package main

import (
	"fmt"
)

func main() {
	mapTest()
}

func mapTest() {
	var m map[string]string
	m = make(map[string]string, 5)
	m["天刀"] = "无名"
	m["太乙"] = "乾坤"
	fmt.Printf("len: %d\n", len(m))
	m["泠月"] = "倚天"
	m["少林"] = "伏魔"
	m["剑王"] = "大夏龙雀"
	m["苍狼"] = "不知道"
	fmt.Printf("len: %d\n", len(m))
	delete(m, "苍狼")
	println("test")
	for key, value := range m {
		fmt.Printf("key: %s\n", key)
		fmt.Printf("value: %s\n", value)
	}

	key := "泠月"
	value, exist := m[key]
	if exist {
		println(value)
	} else {
		fmt.Printf("key %s 不存在value\n", key)
	}
	updateMapValue(m, "太乙", "断江")
	fmt.Printf("after update value: %s", m["太乙"])
}

func updateMapValue(m map[string]string, key string, newValue string) {
	m[key] = newValue
}

func sliceTest() {
	s := make([]int, 0, 5)
	println(s)
	s = append(s, 20)
	println(s)
	for i := 0; i < len(s); i++ {
		println(s[i])
	}
}

func test() {
	var ch chan int
	fmt.Printf("ch is nil %t\n", ch == nil)
	fmt.Printf("len of ch is %d\n", len(ch))

	ch = make(chan int, 10)
	fmt.Printf("len of ch is %d\n", len(ch))
	fmt.Printf("cap of ch is %d\n", cap(ch))

	for i := 0; i < 10; i++ {
		ch <- i
	}
	fmt.Printf("len of ch is %d\n", len(ch))
	fmt.Printf("cap of ch is %d\n", cap(ch))
	_ = <-ch
	_ = <-ch
	ch <- 3

	fmt.Printf("len of ch is %d\n", len(ch))
	fmt.Printf("cap of ch is %d\n", cap(ch))

	close(ch)
	for e := range ch {
		println(e)
	}
}
