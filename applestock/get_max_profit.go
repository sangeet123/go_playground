package applestock

// Integers represents slice of integers
type Integers []int

// Len returns the legth of slice
func (slice Integers) Len() int {
	return len(slice)
}

//GetMaxProfit returns max profit
func GetMaxProfit(stockPrice Integers) int {
	if stockPrice.Len() == 0 {
		return 0
	}

	profit := 0
	purchase := stockPrice[0]

	for _, curPrice := range stockPrice {
		if curProfit := curPrice - purchase; curProfit > profit {
			profit = curProfit
		} else if purchase > curPrice {
			purchase = curPrice
		}
	}
	return profit
}
