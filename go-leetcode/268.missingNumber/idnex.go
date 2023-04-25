package main

import "sort"

func main() {
	nums1 := []int{3, 0, 1}
	nums2 := []int{0, 1}
	nums3 := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}

	println(missingNumber(nums1))
	println(missingNumber(nums2))
	println(missingNumber(nums3))
}

func missingNumber(nums []int) int {
	sort.Ints(nums)

	for i, item := range nums {
		if i != item {
			return i
		}
	}
	return len(nums)
}