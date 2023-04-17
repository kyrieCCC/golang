package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age int `json:"age"`
	Sex string `json:"sex"`
	City string
}

func main() {
	person := Person {
		Name: "wlc",
		Age: 18,
		Sex: "boy",
		City: "hainan",
	}

	//开始序列化操作
	data_json, _ := json.Marshal(&person)

	fmt.Println(person)

	//由于序列化返回的是一个byte形式的切片，我们需要使用string方法来接收字符串类型的值进行输出
	fmt.Println(string(data_json))


	//定义一个新的结构体来接收序列化的结果
	var person_unjson Person

	//将序列化的结果反序列化展示
	err1 := json.Unmarshal([]byte(data_json), &person_unjson)

	if err1 != nil {
		fmt.Println(err1)
	}
	//输出反序列化的结果
	fmt.Println(person_unjson)
}
