package main
import "fmt"

func main(){
	//定义一个切片和map用于遍历
	slice := []int {1,2,3,4,5}
	newMap := map[string]int {
		"a": 1,
		"b": 2,
		"c": 3,
	}

	//开始遍历切片
	for i, val := range slice{
		fmt.Println(i, val)
	}

	fmt.Println("-----遍历map-----")

	//遍历map
	for key, val := range newMap{
		fmt.Println(key, val)
	}
}