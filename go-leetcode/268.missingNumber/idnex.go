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

	println("--------------------------------")
	println(missingNumber_math(nums1))
	println(missingNumber_math(nums2))
	println(missingNumber_math(nums3))
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

//解法三 数学思想
func missingNumber_math(nums []int) int {
	totalSum := 0
	for i := 0; i <= len(nums); i++ {
		totalSum += i
	}

	numsSum := 0
	for _, item := range nums{
		numsSum += item
	}

	return totalSum - numsSum
}