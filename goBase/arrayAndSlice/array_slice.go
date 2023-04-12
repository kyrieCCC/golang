package main
import "fmt"

func main(){
	//初始化一个数组，可以直接指定长度
	//当长度不确定的时候，使用...关键字可以执行判断长度
	//访问某一个下标元素的方式与其他语言类似
	var arr = [...]int{3,2,1}
	a := [5]int{1,2,3,4,5}


	//切片定义
	var slice = make([]int, 3)
	slice2 := [] int {1,2,3,4,5}

	//测试添加与拷贝
	sliceAppend := []int{1,2,3}
	sliceCopy := make([]int, 4)
	sliceCopy = append(sliceCopy, 28)

	fmt.Println(a)
	fmt.Println(arr)
	fmt.Println(slice)
	fmt.Println(slice2)
	
	slice_part(slice2)
	slice_len_cap(slice)
	slice_append_copy(sliceAppend, sliceCopy)
}

//切片截取
func slice_part(slice2 []int){
	slice3 := slice2[1:4]
	fmt.Println(slice3)
}

//切片的len与cap方法
func slice_len_cap(slice []int) {
	fmt.Println(len(slice), cap(slice), slice)
}

//切片的append和copy方法
func slice_append_copy(slice1, slice2 []int) {
	slice1 = append(slice1, 43)
	//这里的slice2我们传一个非空切片便于展示区别
	copy(slice2, slice1)
	fmt.Println(slice1)
	fmt.Println(slice2)
}