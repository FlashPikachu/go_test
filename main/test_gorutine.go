package main

import (
	"sync"
)

var wp = sync.WaitGroup{}

func showMsg(i int) {
	//time.Sleep(10)
	defer wp.Done()
	println(i)
}

func main2() {
	for i := 0; i < 10; i++ {
		go showMsg(i)
		wp.Add(1)
	}

	wp.Wait()

	println("end...")
}
