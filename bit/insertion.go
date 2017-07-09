package bit

func Insert(n int, m int, l, r uint) int {
	var zero uint32 = 0
	zero_comp := ^zero
	left_mask := int(zero_comp << (l + 1))
	right_mask := int(zero_comp >> (32 - r))
	mask := int(left_mask | right_mask)
	toInsert := (mask & n) | (m << r)
	return toInsert
}
