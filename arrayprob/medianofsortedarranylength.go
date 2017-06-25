package arrayprob

import (
	"go_playground/sort"
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

	for skip > 1 && len(arr1) > 0 && len(arr2) > 0 {
		m := compareLengthAndGetMid(len(arr1), len(arr2), skip)
		if arr1[m] >= arr2[m] {
			arr2 = arr2[m:]
		} else {
			arr1 = arr1[m:]
		}
		skip -= m
	}
	arr := append(getFirstTwoElement(arr1), getFirstTwoElement(arr2)...)
	sort.InsertionSortInt(arr)
	return getMedian(arr[skip:], even)

}

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

func getFirstTwoElement(arr []int) []int {
	toReturn := []int{}
	counter := 0
	for _, v := range arr {
		if counter >= 2 {
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
