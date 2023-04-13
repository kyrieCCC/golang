package main
import "fmt"

func main(){
	//定义一个变量和一个指针
	a := 20
	var ip *int

	//将这个指针指向我们的a变量
	ip = &a

	//输出a的存储地址
	fmt.Println("a变量的地址是", &a)

	//输出ip指向的地址，这里因为是指向了a，所以是ip返回的是a的存储地址
	fmt.Println("ip指针指向的地址是", ip)

	//返回ip指向的地址对应的值
	fmt.Println("ip指针指向的地址的值是", *ip)

	add_fail(a)
	fmt.Println(a)
	add2(&a)
	fmt.Println(a)
}


func add2(a *int) {
	*a += 2
}

func add_fail(a int) {
	a += 2
}