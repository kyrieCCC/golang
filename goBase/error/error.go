package main
import (
	"fmt"
	"errors"
)

func main() {
	//定义测试数据
	a := 4

	a, e := sqrt(a)

	//此处进行判断是否有错误信息
	//如果有错误信息输出错误信息，反之输出结果
	if e != nil {
		fmt.Println(e)
		return
	}
	fmt.Println(a)
}

//定义一个计算平方值的函数，具有错误检测
func sqrt(a int)(int, error) {
	if a < 0 {
		return 0, errors.New("negative number!")
	}
	//没有错误，返回nil
	return a*a, nil
}