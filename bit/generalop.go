package bit

func IsSet(n int, pos uint) bool {
	mask := int(1 << pos)
	return (n & mask) > 0
}

func SetBit(n int, pos uint) int {
	mask := int(1 << pos)
	return (n | mask)
}

func GetBit(n int, pos uint) int {
	if IsSet(n, pos) {
		return 1
	}
	return 0
}

func ClearBit(n int, pos uint) int {
	mask := int(^(1 << pos))
	return n & mask
}

// Alternative way to update
// func UpdateBit(n int, pos uint, set bool) int {
// 	toSet := 1
// 	if !set {
// 		toSet = 0
// 	}
// 	mask := int(^(1 << pos))
// 	return int(n & mask  | (toSet << pos))
// }

func UpdateBit(n int, pos uint, set bool) int {
	if set {
		return SetBit(n, pos)
	}
	return ClearBit(n, pos)
}
