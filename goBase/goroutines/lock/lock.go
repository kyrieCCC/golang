package main

import (
	"sync"
	"time"
)

func main() {
	add()
}

//声明一个互斥锁
var (
	x    int64
	lock sync.Mutex
)

//对x进行累加，但是加入锁 
//当对x操作之前先对该协程上锁
//随后对x进行累加
//累加完后释放锁
func addWithLock() {
	for i := 0; i < 2000; i++ {
		lock.Lock()
		x += 1
		lock.Unlock()
	}
}

//同样是对x进行累加，但是此处不加入锁
//不加锁的坏处在于，同时多个协程在跑的时候
//可以同时获取一个状态的x并进行相加，这个时候了累加出来的x就是错误的
func addWithOutLock() {
	for i := 0; i < 2000; i++ {
		x += 1
	}
}

func add(){
	x = 0
	for i := 0; i < 5; i++ {
		//创建五个协程
		//每个协程都是单独进行累加的
		go addWithOutLock()
	}
	//控制时间，不让进程结束
	time.Sleep(time.Second)

	println("withoutLock:", x)

	x = 0
	for i := 0; i < 5; i++ {
		go addWithLock()
	}
	time.Sleep(time.Second)
	println("withLock:", x)
}