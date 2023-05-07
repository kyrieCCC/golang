package main

import "fmt"

func main() {
	test1_1 := []int{1,2,2,1}
	test1_2 := []int{2,2}

	test2_1 := []int{4,9,5}
	test2_2 := []int{9,4,9,8,4}

	fmt.Printf("答案是：%v \n", intersection(test1_1, test1_2))
	fmt.Printf("答案是：%v \n", intersection(test2_1, test2_2))
}

// 349. 两个数组的交集
// 给定两个数组 nums1 和 nums2 ，返回 它们的交集 。
// 输出结果中的每个元素一定是 唯一 的。我们可以 不考虑输出结果的顺序 。
// 输入：nums1 = [1,2,2,1], nums2 = [2,2]
// 输出：[2]
func intersection(nums1 []int, nums2 []int) []int {
	map1 := make(map[int]bool)
	map2 := make(map[int]bool)

	res1 := []int{}
	res2 := []int{}

	for _, item := range nums1 {
		if !map1[item] {
			map1[item] = true
		}
	}

	for _, item := range nums2 {
		if map1[item] {
			res1 = append(res1, item)
		}
	}

	for _, item := range res1 {
		if !map2[item] {
			map2[item] = true
			res2 = append(res2, item)
		}
	}

	return res2
}
//直接使用map来查看两个数组的重合的部分
//再用一个map来给结果数组去重
//得出唯一的值
//最后以数组的形式输出