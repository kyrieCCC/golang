package main

import "sync"

func main() {
	//以结构体的形式声明一个wg
	var wg sync.WaitGroup

	//计数器中添加5
	wg.Add(5)

	for i := 0; i < 5; i++ {
		//创建五个协程
		go func (j int)  {
			//每次协程函数执行完成后都要Done使得计数器减1
			defer wg.Done()
			println(j)
		}(i)
	}

	//等待计数器完成
	wg.Wait()
}