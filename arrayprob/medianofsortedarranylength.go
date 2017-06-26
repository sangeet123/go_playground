package arrayprob

import (
	"go_playground/sort"
)

const (
	threshold = 3
)

func GetMedianOfUnequalLengthArr(arr1, arr2 []int) float32 {
	total := len(arr1) + len(arr2)
	skip := 0
	even := isEven(total)

	if even {
		skip = (total - 1) >> 1
	} else {
		skip = total >> 1
	}

	for skip > threshold && len(arr1) > 0 && len(arr2) > 0 {
		m := compareLengthAndGetMid(len(arr1), len(arr2), skip)
		if arr1[m] >= arr2[m] {
			arr2 = arr2[m:]
		} else {
			arr1 = arr1[m:]
		}
		skip -= m
	}
	arr := mergeArr(getFirstKElement(arr1, threshold), getFirstKElement(arr2, threshold))
	sort.InsertionSortInt(arr)
	return getMedian(arr[skip:], even)

}

// comparison has to be made on equal length chunks
// if mid is greater than skip elements return skip -1
func compareLengthAndGetMid(lenarr1, lenarr2, skip int) int {
	m := lenarr1 >> 1
	if lenarr1 >= lenarr2 {
		m = lenarr2 >> 1
	}

	if m > skip {
		return skip - 1
	}
	return m
}

func getFirstKElement(arr []int, k int) []int {
	toReturn := []int{}
	counter := 0
	for _, v := range arr {
		if counter >= k {
			break
		}
		toReturn = append(toReturn, v)
		counter++
	}
	return toReturn
}

func getMedian(arr []int, even bool) float32 {
	if even {
		return float32(arr[0]+arr[1]) / 2
	}
	return float32(arr[0])
}
