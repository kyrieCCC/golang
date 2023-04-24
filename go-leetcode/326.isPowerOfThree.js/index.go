package main

func main(){

	res1 := isPowerOfThree1(27)
	res2 := isPowerOfThree2(27)

	println(res1)
	println(res2)

}

//解法一 循环
func isPowerOfThree1(n int) bool {
	if n == 1 {
		return true
	}
	for i := 3; i <= n; {
		if i == n {
			return true
		}
		i = i * 3
	}
	return false
}

//解法二 试除法
func isPowerOfThree2(n int) bool {
	for n > 0 && n % 3 == 0 {
		n /= 3
	}
	return n == 1
}

