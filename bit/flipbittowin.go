package bit

// The only case it fails is when everything is one
func GetLongFlipLen(no int) int {
	// initialization
	msb := GetMSBSetBit(no)
	right := 0
	left := 0
	longest := 0

	for i := 0; i <= msb+1; {
		if !IsSet(no, uint(i)) {
			j := i + 1
			for ; IsSet(no, uint(j)); j++ {
				left++
			}
			if longest < (right + left + 1) {
				longest = right + left + 1
			}
			right = left
			left = 0
			i = j
		} else {
			right++
			i++
		}
	}
	return longest
}

func GetMSBSetBit(no int) int {
	i := -1
	if no < 0 {
		no *= -1
	}
	for ; no != 0; i++ {
		no = no >> 1
	}
	return i
}
