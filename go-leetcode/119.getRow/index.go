package main

import "fmt"

func main() {
	res1 := getRow(3)
	res2 := getRow(2)
	fmt.Printf("结果为: %v \n", res1)
	fmt.Printf("结果为: %v \n", res2)
}


func getRow(rowIndex int) []int {
	res := make([][]int, rowIndex + 1)
	for i := range res {
		res[i] = make([]int, i + 1)
		res[i][0], res[i][i] = 1, 1
		for j := 1; j < i; j++ {
			res[i][j] = res[i - 1][j - 1] + res[i - 1][j]
		}
	}
	return res[rowIndex]
}
//直接使用递推的方式
//计算出每一个位置应该的值
//最后返回指定的rowIndex那一行的值即可