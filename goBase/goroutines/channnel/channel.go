package main


//测试协程间的通信
//主要步骤为
//首先创建两个 channel
//A channel中向B channel传递 1-9 的数字
//B channel首先接收数据，然后将数字进行平方的计算后再放到B channel
//最后读取B channel中的数据，打印出结果

func main() {
	//创建两个channel
	send := make(chan int)
	recive := make(chan int, 3)

	//将1 - 9 的数字丢入A channel中
	go func(){
		//每次操作结束后关闭channel
		defer close(send)
		for i := 0; i < 10; i++ {
			send <- i
		}
	}()

	//读取A channel中的数据，并进行平方运算后丢入B channel
	go func() {
		defer close(recive)
		for i := range send {
			recive <- i * i
		}
	}()

	//读取B channel中的数字并打印
	for i := range recive {
		println(i)
	}
}