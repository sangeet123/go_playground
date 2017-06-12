package coinexchange

import "sort"

//Integers is integers slice
type Integers []int

//Len returns the size of slice
func (slice Integers) Len() int {
	return len(slice)
}

//GetTotalWays retruns total ways to make coins
func GetTotalWays(total int, coins Integers) int {

	n := coins.Len()
	if total == 0 || n == 0 {
		return 0
	}

	m := total + 1

	//memoization
	storage := make(Integers, m)
	storage[0] = 1

	//sorting coins
	sort.Ints(coins)

	for i := 1; i < m; i++ {
		ii := i - coins[0]
		if ii >= 0 {
			storage[i] = storage[ii]
		}
	}

	for i := 1; i < n; i++ {
		coin := coins[i]
		for j := 1; j < m; j++ {
			jj := j - coin
			if jj >= 0 {
				storage[j] += storage[jj]
			}
		}
	}

	return storage[len(storage)-1]
}
