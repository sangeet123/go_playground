package arrayprob

import (
	"go_playground/sort"
)

// no validation of length done on array
// equal length array should be provided
// also array of zero length are not expected
//GetMedian returns lower median
func GetMedian(arr1, arr2 []int) int {

	for len(arr1) > 1 && len(arr2) > 1 {
		m1 := (len(arr1) - 1) >> 1
		m2 := (len(arr2) - 1) >> 1

		if arr1[m1] == arr2[m2] {
			return arr1[m1]
		}

		if arr1[m1] > arr2[m2] {
			arr1 = arr1[0 : m1+1]
			if len(arr2)%2 == 0 {
				m2++
			}
			arr2 = arr2[m2:]
		} else {
			if len(arr1)%2 == 0 {
				m1++
			}
			arr1 = arr1[m1:]
			arr2 = arr2[0 : m2+1]
		}
	}

	arr := mergeArr(arr1, arr2)
	sort.InsertionSortInt(arr)
	m := (len(arr) - 1) >> 1
	return arr[m]
}

// utility for merging two array
func mergeArr(arr1, arr2 []int) []int {
	if len(arr1) < len(arr2) {
		return mergeArr(arr2, arr1)
	}
	arr := []int{}
	for _, v := range arr1 {
		arr = append(arr, v)
	}

	for _, v := range arr2 {
		arr = append(arr, v)
	}
	return arr
}
