package bit

func Conversion(a int, b int) int {
	return CountBitSet(a ^ b)
}
