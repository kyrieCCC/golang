package main

import (
	"sort"
	"fmt"
)

func main() {
	nums1 := []int{3, 2, 1}
	nums2 := []int{2, 1}
	nums3 := []int{3, 2, 2, 1}

	fmt.Printf("答案是: %v \n", thirdMax(nums1))
	fmt.Printf("答案是: %v \n", thirdMax(nums2))
	fmt.Printf("答案是: %v \n", thirdMax(nums3))
}


func thirdMax(nums []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
  
	for i, diff := 1, 1; i < len(nums); i++ {
	  if nums[i - 1] != nums[i] {
		diff++
		if diff == 3 {
		  return nums[i]
		}
	  }
	}
	return nums[0]
  }