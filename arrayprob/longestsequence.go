package arrayprob

func LengthOfLongestSequence(arr []int) int {
	numbersMap := make(map[int]struct{})

	for _, v := range arr {
		numbersMap[v] = struct{}{}
	}

	longestLength := 0
	for i := 0; i < len(arr) && len(numbersMap) > 0; i++ {
		currentLongest := getLongestSequence(numbersMap, arr[i], 1) +
			getLongestSequence(numbersMap, arr[i]-1, -1)
		if currentLongest > longestLength {
			longestLength = currentLongest
		}
	}
	return longestLength
}

func getLongestSequence(numbersMap map[int]struct{}, no int, step int) int {
	_, isPresent := numbersMap[no]
	currentLongest := 0
	for isPresent {
		currentLongest++
		delete(numbersMap, no)
		no += step
		_, isPresent = numbersMap[no]
	}
	return currentLongest
}
