package arrayprob

func GetInsertionPos(arr []int, no int) int {
	l := 0
	r := len(arr) - 1
	for l <= r {
		m := (l + r) >> 1
		if arr[m] == no {
			return m
		} else if arr[m] > no {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return l
}
