package sort

import "math/rand"

func QuickSortInt(nos integers) {
	quickSortInt(0, len(nos)-1, nos)
}

//randomized quick sort
func quickSortInt(l int, r int, nos integers) {
	if l < r {

		// randomizing quick sort so that sorted array works with
		// expected run time of O(nlogn)
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

		quickSortInt(l, ii-1, nos)
		quickSortInt(ii+1, r, nos)
	}
}

func getRandomIntInRange(min, max int) int {
	return rand.Intn(max-min) + min
}
