package arrayprob

func GetMinSizeSubArr(arr []int, sum int) []int {
	if len(arr) == 0 {
		return []int{}
	}

	//initialization
	ri, rj, i, j := 0, len(arr), 0, 1
	csum := arr[0]
	isPresent := false

	//This is done to take care of last element of array
	dummyEndArr := append(arr, 0)
	for j < len(dummyEndArr) {
		if csum < sum {
			csum += dummyEndArr[j]
			j++
		}
		if csum >= sum {
			isPresent = true
			if j-i < rj-ri {
				rj, ri = j, i
			}
			csum -= dummyEndArr[i]
			i++
			if i == j {
				csum = dummyEndArr[i]
				j++
			}
		}
	}
	// if no such subarray present return empty array
	if !isPresent {
		return []int{}
	}

	return arr[ri:rj]
}
