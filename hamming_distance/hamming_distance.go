package hammingdistance

// returns how many bits are set on a number
func GetNoOfOnes(no int) int {
	if no == 0 {
		return 0
	}

	return 1 + GetNoOfOnes(no&(no-1))
}

// return hamming distance
func GetHammingDistance(a int, b int) int {
	// xor operation between sets those bit which are different
	return GetNoOfOnes(a ^ b)
}
