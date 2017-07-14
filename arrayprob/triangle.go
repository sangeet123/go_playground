package arrayprob

func GetMin(arr [][]int) int {
	r := make([]int, len(arr)+1)
	for i := len(arr) - 1; i >= 0; i-- {
		for j := 0; j < len(arr[i]); j++ {
			r[j] = min(arr[i][j]+r[j], arr[i][j]+r[j+1])
		}
	}
	return r[0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
