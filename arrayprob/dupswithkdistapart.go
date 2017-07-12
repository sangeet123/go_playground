package arrayprob

func HasDups(arr []int, k int) bool {
	tracer := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		if ind, ok := tracer[arr[i]]; !ok {
			tracer[arr[i]] = i
		} else if (i - ind) >= k {
			return true
		}
	}
	return false
}
