package bit

func CountBitSet(no int) int {
	if no == 0 {
		return 0
	}
	return CountBitSet(no&(no-1)) + 1
}
