package main

func main(){
	res1 := isPowerOfTwo1(5)

	res2 := isPowerOfTwo3(123)

	res3 := isPowerOfTwo2(256)
	
	println(res1)
	println(res2)
	println(res3)
}

//1. 循环
func isPowerOfTwo1(n int) bool {
	if n == 1 {
		return true
	}
	// val := 2
	for val := 2; val <= n;{
		if n == val {
			return true
		}
		val = val * 2
	}
	return false
}

//2. 递归
func isPowerOfTwo2(n int) bool {
	if n == 1 {
		return true
	}
	if n > 2 && n != 1 {
		return false
	}
	return isPowerOfTwo2(n / 2)
}

//3. 位运算
func isPowerOfTwo3(n int) bool {
	return n > 0 && n & (n - 1) == 0
}
// 击败100% ！！！！
