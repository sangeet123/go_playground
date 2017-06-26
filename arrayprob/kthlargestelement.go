package arrayprob

import (
	"math/rand"
)

func GetKthLargestElement(nos []int, k int) int {
	toReturn := 0
	for k > 0 {

		// exact copy of quick sort
		// removing slice from an array that is not required
		// is the mechanism and avoiding recursion
		l := 0
		r := len(nos) - 1
		randIndex := getRandomIntInRange(l, r)
		nos[r], nos[randIndex] = nos[randIndex], nos[r]

		pivot := nos[r]
		ii := l - 1
		for i := l; i < r; i++ {
			if nos[i] < pivot {
				ii++
				nos[ii], nos[i] = nos[i], nos[ii]
			}
		}

		ii++
		nos[ii], nos[r] = nos[r], nos[ii]
		if (k - 1) == ii {
			toReturn = nos[ii]
			break
		}

		if (k - 1) > ii {
			nos = nos[ii+1:]
			k = k - ii - 1
		} else {
			nos = nos[0:ii]
		}
	}
	return toReturn
}

func getRandomIntInRange(min, max int) int {
	if min == max {
		return min
	}
	return rand.Intn(max-min) + min
}
