package main
import "fmt"

func main() {
	a := 100
	b := 200

	//因为返回两个参数，所以需要两个值来接收结果
	a, b = swap(a, b)

	fmt.Println(a, b)
	//返回200 100 表示成功交换两个数的值
}

//定义一个新的函数，交换两个数的值，并且返回两个参数
func swap(p1, p2 int)(int, int){
	return p2, p1
}