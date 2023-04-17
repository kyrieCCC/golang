// 1. 两数之和 
// 给定一个整数数组 nums 和一个整数目标值 target，
// 请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
// 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
// 你可以按任意顺序返回答案。

// 输入：nums = [2,7,11,15], target = 9
// 输出：[0,1]
// 解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 

package main
import "fmt"

func main()  {
	nums1 := []int{2,7,11,15}
	target1 := 9

	resSlice := toSum(nums1, target1)

	fmt.Println(resSlice)
}

func toSum(nums []int, target int) []int {
	for i, val := range nums {
		for j := i + 1; j < len(nums); j++ {
			if val + nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}