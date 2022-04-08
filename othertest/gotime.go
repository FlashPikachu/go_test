package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var i int32 = 100
var lock = sync.Mutex{}

func output() {
	for i := 0; i < 2; i++ {
		//time.Sleep(time.Second)
		println(i)
	}
}

func add() {
	atomic.AddInt32(&i, 1)

	//lock.Lock()
	//i += 1
	//lock.Unlock()
	//fmt.Printf("i++: %d\n", i)
}

func sub() {
	atomic.AddInt32(&i, -1)
	//lock.Lock()
	//i -= 1
	//lock.Unlock()
	//fmt.Printf("i--: %d\n", i)
}

func main() {
	for i := 0; i < 100; i++ {
		go add()
		go sub()
	}
	time.Sleep(time.Second)
	fmt.Printf("i: %d", i)
	//runtime.GOMAXPROCS(1)

	//runtime.Goexit()
	//go output()
	//for i := 0; i < 3; i++ {
	//	runtime.Gosched()
	//	fmt.Printf("test %d\n", i)
	//}
	//fmt.Printf("cpu core num: %d", runtime.NumCPU())
	//time.Sleep(time.Second)
	//println("Test")
	//test()
}

func test() {
	timer := time.NewTimer(time.Second * 3)
	fmt.Printf("现在的时间是 %s\n", time.Now())
	//x := <-timer.C
	<-timer.C
	fmt.Printf("现在的时间是 %s", time.Now())
}
