package main

import "sort"

func main() {
	nums1 := []int{1, 2, 3, 1}
	nums2 := []int{1, 2, 3, 4}

	println(containsDuplicate(nums1))
	println(containsDuplicate(nums2))

	println("-------------------------------------------")
	println(containsDuplicate_sort(nums1))
	println(containsDuplicate_sort(nums2))
}

// 217. 存在重复元素
// 给你一个整数数组 nums 。如果任一值在数组中出现 至少两次 ，返回 true ；
// 如果数组中每个元素互不相同，返回 false 。
// 输入：nums = [1,2,3,1]
// 输出：true

// 解法一 哈希表
func containsDuplicate(nums []int) bool {
	myMap := make(map[int]bool, 0)

	for _, item := range nums {
		if myMap[item] {
			return true
		} else {
			myMap[item] = true
		}
	}
	return false
}

// 解法二 排序
func containsDuplicate_sort(nums []int) bool {
	sort.Ints(nums)

	for i := 0; i < len(nums) - 1; i++ {
		if nums[i] == nums[i + 1] {
			return true
		}
	} 
	return false
}