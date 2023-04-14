package main

import (
	"fmt"
	"strings"
)

func main(){
	//定义一个字符串
	a := "hello world"

	//判断是否有指定内容
	fmt.Println(strings.Contains(a, "h")) //true
	//判断指定字符出现的次数
	fmt.Println(strings.Count(a, "o")) //2
	//判断指定的前缀
	fmt.Println(strings.HasPrefix(a, "he"))//true
	//判断指定的后缀
	fmt.Println(strings.HasSuffix(a, "ld"))//true
	//判断字符第一次出现的下标
	fmt.Println(strings.Index(a, "o"))//4
	//拼接字符串
	fmt.Println(strings.Join([]string{"wo", "zai", "meituan"}, "-"))//wo-zai-meituan
	//重复叠加字符串
	fmt.Println(strings.Repeat(a, 2))//hello worldhello world
	
	b := "phzyzcr"

	//元素替换
	fmt.Println(strings.Replace(b, "zyzcr", "!", -1))//ph!
	//元素分割为切片形式
	fmt.Println(strings.Split(b, ""))//[p h z y z c r]
	//转换为小写
	fmt.Println(strings.ToLower(b))//phzyzcr
	//转换为大写
	fmt.Println(strings.ToUpper(b))//PHZYZCR
}