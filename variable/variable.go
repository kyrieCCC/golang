package main

import "fmt"

func main() {
   /* 定义局部变量 */
   var a int = 300
   var b int = 200
   var c int = 150
   var ret int
   var ans int

   /* 调用函数并返回最大值 */
   ret = max(a, b)
   ans = min(a, b, c)
   fmt.Printf( "最大值是 : %d\n", ret )
   fmt.Printf( "最小值是 : %d\n", ans )
}

/* 函数返回两个数的最大值 */
func max(num1, num2 int) int {
   /* 定义局部变量 */
   var result int

   if (num1 > num2) {
      result = num1
   } else {
      result = num2
   }
   return result
}

func min(num1, num2, num3 int) int {
   var res int
   if (num1 < num2 && num2 > num3) {
      res = num1
   }
   if (num2 < num1 && num1 > num3) {
      res = num2
   }
   if (num3 < num1 && num1 > num2) {
      res = num3
   } else {
      res = 0
   }
   return res
}