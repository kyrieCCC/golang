package main

import (
	"sync"
	"time"
)

func main() {
	add()
}

var (
	x    int64
	lock sync.Mutex
)

func addWithLock() {
	for i := 0; i < 2000; i++ {
		lock.Lock()
		x += 1
		lock.Unlock()
	}
}

func addWithOutLock() {
	for i := 0; i < 2000; i++ {
		x += 1
	}
}

func add(){
	x = 0
	for i := 0; i < 5; i++ {
		go addWithOutLock()
	}
	time.Sleep(time.Second)

	println("withoutLock:", x)

	x = 0
	for i := 0; i < 5; i++ {
		go addWithLock()
	}
	time.Sleep(time.Second)
	println("withLock:", x)
}