package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := 12
	a1 := "12345"
	a2 := "3.14"
	num := 3.43

	//转换为字符串
	b := strconv.Itoa(a)

	//整数转浮点型
	f := float64(a)

	//字符串转整型
	i, _ := strconv.Atoi(a1)

	//字符串转浮点型
	f2, _ := strconv.ParseFloat(a2, 64)

	//浮点型转字符串
	i2 := strconv.FormatFloat(num, 'f', 2, 64)
	
	fmt.Println(b)
	fmt.Println(f)
	fmt.Println(i)
	fmt.Println(f2)
	fmt.Println(i2)
}