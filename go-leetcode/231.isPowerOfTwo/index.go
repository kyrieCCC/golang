package main

func main(){
	res1 := isPowerOfTwo1(5)

	println(res1)
}

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