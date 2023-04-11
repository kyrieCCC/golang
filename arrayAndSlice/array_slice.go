package main
import "fmt"

func main(){
	//初始化一个数组，可以直接指定长度
	//当长度不确定的时候，使用...关键字可以执行判断长度
	//访问某一个下标元素的方式与其他语言类似
	var arr = [...]int{3,2,1}
	a := [5]int{1,2,3,4,5}


	//切片定义
	var slice = make([]int, 0)
	slice2 := [] int {1,2,3,4,5}

	fmt.Println(a)
	fmt.Println(arr)
	fmt.Println(slice)
	fmt.Println(slice2)
	
	slice_part(slice2)
}

//切片截取
func slice_part(slice2 []int){
	slice3 := slice2[1:4]
	fmt.Println(slice3)
}