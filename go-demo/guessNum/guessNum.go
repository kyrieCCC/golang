package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)


//实现一个简单的猜字游戏
func main() {
	maxNum := 100

	//生成一个随机数，使用rand.Intn方法
	randomNum := rand.Intn(maxNum)

	// fmt.Println(randomNum)

	fmt.Println("please input your number:")

	//进行循环猜谜
	for {
		//创建一个新的可读流，并读取命令台输入的内容
		reader := bufio.NewReader(os.Stdin)

		//获取到输入的内容
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("error1", err)
			// return
			continue
		}

		//去除换行符
		input = strings.TrimSuffix(input, "\r\n")

		//将输入的内容转换为int整型
		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("error2", err)
			// return
			continue
		}
		// fmt.Println("your guess number is:", guess)


		//不断循环判断输入的数字大小，直到输入的结果正确
		if guess < randomNum {
			fmt.Println("a little small")
		} else if guess > randomNum {
			fmt.Println("a little big")
		} else {
			fmt.Println("right, this num is :", randomNum)
			break
		}
	}

}
