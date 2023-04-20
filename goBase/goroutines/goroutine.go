package main

import (
	"fmt"
	"time"
)

func main() {
	//一个for循环，每次执行的时候都使用go关键字创建一个协程
	for i := 0; i < 5; i++ {
		// 创建一个协程
		go printNum(i)
	}

	//阻塞时间，保证在协程内容执行完前主线程不退出
	time.Sleep(time.Second)
}

func printNum(j int) {
	fmt.Println(j)
}