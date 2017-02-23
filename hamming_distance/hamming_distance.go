package hammingdistance

func GetNoOfOnes(no int) int {
		if no == 0 {
			return 0
		}

		return 1 + GetNoOfOnes(no & (no - 1))	
}

func GetHammingDistance(a int, b int) int {
	return GetNoOfOnes(a^b)
}