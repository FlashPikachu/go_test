package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

var path string = "E:/temp/test.txt"

func main() {
	test8()
}

func test1() {
	//只读的方式打开path目录下的文件
	file, err := os.Open(path)
	if err == nil {
		fmt.Println("open file successfully")
	} else {
		fmt.Printf("failed to open file , err=%v\n", err)
		return
	}
	defer file.Close()
	fmt.Println("file:", file)
}

func test2() {
	//读取文件
	//file.Read()
	file, err := os.Open(path)
	if err == nil {
		fmt.Println("open file successfully")
	} else {
		fmt.Printf("failed to open file , err=%v\n", err)
		return
	}
	defer file.Close()
	var tmp = make([]byte, 128)
	n, err := file.Read(tmp)
	if err != nil {
		fmt.Printf("read from file failed err: %v\n", err)
		return
	} else {
		fmt.Printf("read %d bytes from file\n", n)
		fmt.Println(string(tmp))
	}

}

func test3() {
	//循环读取文件
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("failed to open file , err=%v\n", err)
		return
	}

	defer file.Close()

	for {
		var tmp = make([]byte, 4)
		n, err := file.Read(tmp)
		if err == io.EOF {
			fmt.Println(string(tmp[:n]))
			return
		}
		if err != nil {
			fmt.Printf("failed to open file , err=%v\n", err)
			return
		}
		fmt.Printf("read %d bytes from file\n", n)
		fmt.Println(string(tmp))
	}
}

func test4() {
	//bufio 读取文件
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("failed to open file , err=%v\n", err)
		return
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	for true {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println(line)
			return
		}
		if err != nil {
			fmt.Printf("read from file failed err: %v\n", err)
			return
		}
		fmt.Print(line)
	}

}

func test5() {
	//ioutil读取整个文件
	file, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("failed to open file , err=%v\n", err)
		return
	}
	fmt.Println("test")
	fmt.Println(string(file))
}

func test6() {
	//文件写入操作
	file, err := os.OpenFile(path, os.O_APPEND, 0755)
	if err != nil {
		fmt.Printf("failed to open file , err=%v\n", err)
		return
	}
	defer file.Close()
	str := "601919 yes!"
	file.Write([]byte(str))
	str = "gogogo"
	file.WriteString(str)
}

func test7() {
	//使用bufio写
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC, 0755)
	if err != nil {
		fmt.Printf("failed to open file , err=%v\n", err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	writer.WriteString("test11111")
	//将缓冲区的内容写入文件
	writer.Flush()
}

func test8() {
	//使用ioutil写
	str := "勇敢的海斗士，中远海皇！！"
	err := ioutil.WriteFile(path, []byte(str), 0755)
	if err != nil {
		fmt.Printf("write file err :%v", err)
		return
	}

}

func CopyFile(source string, destination string) {
	//todo
}
