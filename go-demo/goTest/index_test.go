package main

import (
	"testing"
)

//测试文件
//下面的函数为测试函数
//可以用于测试index.go中的指定函数是否可以通过测试
func TestHelloWlc(t *testing.T) {
	output := outputWlc()
	//可以通过
	expectOutPut := "Wlc"


	//不可以通过
	// expectOutPut1 := "Tom"


	//如果不通过的报错信息
	if output != expectOutPut {
		t.Errorf("Expected do not match")
	}
}