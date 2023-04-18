// 111. 二叉树的最小深度
// 给定一个二叉树，找出其最小深度。
// 最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
// 说明：叶子节点是指没有子节点的节点。
// 输入：root = [3,9,20,null,null,15,7]
// 输出：2

package main

import (
	"fmt"
	"math"
)
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	var subNode TreeNode 

	subNode.Val = 2
	subNode.Left = &subNode
	subNode.Right = &subNode
	ans := minDepth(&subNode)
	fmt.Println(ans)
}



func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}

	res := math.MaxInt32

	if root.Left != nil {
		res = min(minDepth(root.Left), res)
	}
	if root.Right != nil {
		res = min(minDepth(root.Right), res)
	}
	return res + 1
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
