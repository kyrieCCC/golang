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
	resSlice2 := toSum2(nums1, target1)

	fmt.Println(resSlice)
	fmt.Println(resSlice2)
}

//两数之和解法一，暴力解法
func toSum(nums []int, target int) []int {
	//直接通过双循环来进行遍历
	//找出确定先行值与其次值来确定两数的和
	for i, val := range nums {
		for j := i + 1; j < len(nums); j++ {
			if val + nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

//两数之和解法二，哈希表
func toSum2(nums []int, target int) []int {
	newMap := map[int]int{}

	for i, val := range nums {
		if j, ok := newMap[target - val]; ok {
			return []int{j, i}
		}
		newMap[val] = i
	}
	return nil
}