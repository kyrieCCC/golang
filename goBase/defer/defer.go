package main

import "fmt"

func main() {
	//不添加defer来执行
	for i := 0; i < 5; i++ {
		fmt.Print(i, " ") //执行结果为0 1 2 3 4
	}

	fmt.Println("")

	//添加defer参数
	for j := 0; j < 5; j++ {
		defer fmt.Print(j, " ") //执行结果为4 3 2 1 0
	}
	//执行结果的不同是由于defer关键字为延迟执行，执行顺序变为LIFO后进先出
}