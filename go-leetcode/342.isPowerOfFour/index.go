package main


func main() {
	a1 := 1
	a2 := 5
	a3 := 64

	println(isPowerOfFour(a1))
	println(isPowerOfFour(a2))
	println(isPowerOfFour(a3))
	println("=====================================")
	println(isPowerOfFour_deep(a1))
	println(isPowerOfFour_deep(a2))
	println(isPowerOfFour_deep(a3))
}

//解法一 循环
func isPowerOfFour(n int) bool {
	if n == 1 {
		return true
	}
	for i := 4; i <= n; {
		if i == n {
			return true
		}
		i = i * 4
	}
	return false
}

//解法二 递归
func isPowerOfFour_deep(n int) bool {
	return n > 0 && n & (n - 1) == 0 && n % 3 == 1
}