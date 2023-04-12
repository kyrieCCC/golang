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
	// v1 := newMap["name"]
	// fmt.Println(v1)

	addVal(newMap)
	getLen(newMap)
	getAllMap(newMap)
	deleteVal(newMap)
}

//添加键值对，key为num, val为43
func addVal(newMap map[string]string) {
	newMap["num"] = "43"
	fmt.Println(newMap)
}

//获取长度
func getLen(newMap map[string]string) {
	len := len(newMap)
	fmt.Println("map的长度为", len)
}

//遍历map
func getAllMap(newMap map[string]string) {
	for k, val := range newMap  {
		fmt.Println(k, val)
	}
}

//删除键值对
func deleteVal(newMap map[string]string) {
	delete(newMap, "num")
	fmt.Println(newMap)
}