package arrayprob

//TwoSumArr for sorted array
func TwoSumArr(arr []int, sum int) (bool, int, int) {
	for i, j := 0, len(arr)-1; i < j; {
		if arr[i]+arr[j] == sum {
			return true, i, j
		} else if arr[i]+arr[j] > sum {
			j--
		} else {
			i++
		}
	}
	return false, -1, -1
}

// Returns immediately the first one found
// Can be made to return other elements as array list
//FourSumArr for sorted array
func FourSumArr(arr []int, sum int) (bool, []int) {
	for i := 0; i < len(arr)-3; i++ {
		for j := i + 1; j < len(arr)-2; j++ {
			if hasSum, k, l := TwoSumArr(arr[j+1:], sum-arr[i]-arr[j]); hasSum {
				return true, []int{i, j, j + 1 + k, j + 1 + l}
			}
		}
	}
	return false, []int{-1, -1, -1, -1}
}
