package main

import (
	"fmt"
	"sync"
	"time"
)

// make 和 new都是用来初始化内存
// new多用来初始化基本数据类型，返回的是指针 初始化bool string int...
// make返回的是对应类型 初始化slice map channel

//how to let goroutine exit
//1 全局变量
var exit bool

//2 channel

var wg sync.WaitGroup

func worker(ch <-chan bool) {
	defer wg.Done()
LABEL:
	for true {
		//if exit {
		//	break
		//}
		select {
		case <-ch:
			break LABEL
		default:
			fmt.Println("worker...")
			time.Sleep(time.Second)
		}
	}
}

func main() {

	//wg.Add(1)
	//go worker()
	//time.Sleep(time.Second * 3)
	//exit = true //通过修改全局变量使子goroutine退出
	//wg.Wait()

	var exitChan = make(chan bool)
	wg.Add(1)
	go worker(exitChan)
	exitChan <- true
	wg.Wait()

	fmt.Println("work over!!!")

}
