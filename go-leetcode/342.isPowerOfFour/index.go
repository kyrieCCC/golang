package main


func main() {
	a1 := 1
	a2 := 5
	a3 := 64

	println(isPowerOfFour(a1))
	println(isPowerOfFour(a2))
	println(isPowerOfFour(a3))
}

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