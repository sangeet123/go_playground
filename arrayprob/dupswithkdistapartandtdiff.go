package arrayprob

import (
	"go_playground/sort"
)

func Exists(arr []int, t, k int) bool {

	if len(arr) == 0 {
		return false
	}

	if k == 0 && t <= 0 {
		return true
	}

	if k <= 0 {
		return false
	}

	if k >= len(arr) {
		k = len(arr) - 1
	}

	// O(k)
	kelements := copyOf(arr, 1, k+1)

	// O(klogk)
	sort.QuickSortInt(kelements)

	// O(n-k)
	for i := 0; i < len(arr)-1; i++ {
		if abs(arr[i]-kelements[0]) <= t {
			return true
		}

		// log(k)
		// removes the first element
		pos := GetInsertionPos(kelements, arr[i+1])
		kelements = append(kelements[0:pos], kelements[pos+1:]...)

		// log(k)
		// append the next element
		if k+i+1 < len(arr) {
			pos := GetInsertionPos(kelements, arr[k+i+1])
			kelements = append(append(kelements[0:pos], arr[k+i+1]), kelements[pos:]...)
		}
	}
	return false
}

func copyOf(src []int, start, end int) []int {
	result := []int{}
	for i := start; i < end; i++ {
		result = append(result, src[i])
	}
	return result
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
