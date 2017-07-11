package bit

// Preserves sign and the 31st bit as it does not have pair to swap
func SwapEvenOddBit(no int32) int32 {
	var evenMask int32 = 0x15555555
	var oddMask int32 = 0x2AAAAAAA
	var signAndLastBitMask int32 = -1073741824
	return no&evenMask<<1 | no&oddMask>>1 | no&signAndLastBitMask
}
