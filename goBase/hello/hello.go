// 定义了一个包名，必须在源文件的第一行非注释来说明文件属于哪个包
// 每一个go程序都需要包含一个名为main的包
package main

// import "fmt" 告诉 Go 编译器这个程序需要使用 fmt 包（的函数，或其他元素）
// fmt 包实现了格式化 IO（输入/输出）的函数。
import (
	"fmt"
	"strings"
)

func main(){
	//换行
	fmt.Println("hello world")

	fmt.Print("hello world")

	slices := []string {"hello", "world"}

	//通过切片合并的方式将一个有指定元素的切片合成为一个字符串
	outHello((slices))
}

func outHello(slice []string) {
	//join合并操作
	fmt.Println(strings.Join(slice, "-"))
}
