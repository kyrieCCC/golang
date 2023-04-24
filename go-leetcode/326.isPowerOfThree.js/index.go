package main

func main(){

	res1 := isPowerOfThree1(27)

	println(res1)
}

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