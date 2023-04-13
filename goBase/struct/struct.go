package main
import "fmt"


//定义一个结构体
type newStruct struct {
	name string
	age int
	company string
}


func main(){
	//创建一个实例
	var testStruct newStruct

	//在主函数中为结构体进行赋值操作
	testStruct.name = "wlc"
	testStruct.age = 18
	testStruct.company = "meituan"

	//输出查看
	fmt.Println(testStruct.name)
	fmt.Println(testStruct.age)
	fmt.Println(testStruct.company)

	//传入一个结构体的地址
	changeVal(&testStruct)
}

//通过指针来修改源结构体里面的值
func changeVal(changeStruct *newStruct)  {
	changeStruct.age = 20
	fmt.Println(changeStruct.age)
}