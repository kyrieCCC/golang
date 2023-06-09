// 344. 反转字符串
// 编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组 s 的形式给出。
// 不要给另外的数组分配额外的空间，你必须原地修改输入数组、使用 O(1) 的额外空间解决这一问题。
// 输入：s = ["h","e","l","l","o"]
// 输出：["o","l","l","e","h"]
package main

import "fmt"

func reverseString(s []string)  {
	for left, right :=0, len(s) - 1; left <= right; left++ {
		s[left], s[right] = s[right], s[left]
		right--
	}
}
//击败99.8%
//双指针解法，创建两个指针分别指向前端与后端
//然后两两交换位置
//终止条件为两个指针的位置重合，代表交换结束

func main() {
	s := []string{"h","e","l","l","o"}
	reverseString(s)
	fmt.Println(s)
}