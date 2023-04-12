package main
import "fmt"

func main(){
	//声明一个map
	newMap := map[string]string {
		"name": "wlc",
		"age": "18",
		"company": "meituan",
	}

	//通过key来获取某一个value
	v1 := newMap["name"]
	fmt.Println(v1)
}