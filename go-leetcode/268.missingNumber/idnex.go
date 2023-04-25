package main

import "sort"

func main() {
	nums1 := []int{3, 0, 1}
	nums2 := []int{0, 1}
	nums3 := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}

	println(missingNumber(nums1))
	println(missingNumber(nums2))
	println(missingNumber(nums3))

	println("--------------------------------")
	println(missingNumber_map(nums1))
	println(missingNumber_map(nums2))
	println(missingNumber_map(nums3))
}

//解法一 暴力循环
func missingNumber(nums []int) int {
	sort.Ints(nums)

	for i, item := range nums {
		if i != item {
			return i
		}
	}
	return len(nums)
}

//解法二 哈希
func missingNumber_map(nums []int) int {
	myMap := map[int]bool{}

	for _, item := range nums {
		myMap[item] = true
	}
	for i := 0; ;i++ {
		if !myMap[i] {
			return i
		}
	}
}
