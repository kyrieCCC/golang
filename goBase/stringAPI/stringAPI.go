package main

import (
	"fmt"
	"strings"
)

func main(){
	//定义一个字符串
	a := "hello world"

	fmt.Println(strings.Contains(a, "h")) //true
	fmt.Println(strings.Count(a, "o")) //2
	fmt.Println(strings.HasPrefix(a, "he"))//true
	fmt.Println(strings.HasSuffix(a, "ld"))//true
	fmt.Println(strings.Index(a, "o"))//4
	fmt.Println(strings.Join([]string{"wo", "zai", "meituan"}, "-"))//wo-zai-meituan
	fmt.Println(strings.Repeat(a, 2))//hello worldhello world
	
	b := "phzyzcr"
	fmt.Println(strings.Replace(b, "zyzcr", "!", -1))//ph!
	fmt.Println(strings.Split(b, ""))//[p h z y z c r]
	fmt.Println(strings.ToLower(b))//phzyzcr
	fmt.Println(strings.ToUpper(b))//PHZYZCR
}