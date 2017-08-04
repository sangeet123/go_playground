package math

func Pow(n, x int) int {
	if x == 0 {
		return 1
	}
	result := Pow(n, x/2)
	result *= result
	if IsOdd(n) {
		result *= x
	}
	return result
}

func IsOdd(n) bool {
	return n%2 != 0
}
