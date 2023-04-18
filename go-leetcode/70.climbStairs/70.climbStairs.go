// 70. 爬楼梯
// 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
// 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
// 输入：n = 2
// 输出：2
// 解释：有两种方法可以爬到楼顶。
// 1. 1 阶 + 1 阶
// 2. 2 阶

package main
import "fmt"


func climbStairs(n int) {
	q, p, res := 0, 0, 1
	
	for i := 1; i <= n; i++ {
		p = q
		q = res
		res = p + q
	}
	fmt.Println(res)
}
