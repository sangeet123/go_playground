package applestock

type Integers []int

func (slice Integers) Len() int {
	return len(slice)
}

func GetMaxProfit(stockPrice Integers) int {
	if stockPrice.Len() == 0 {
		return 0
	}

	profit := 0
	purchase := stockPrice[0]

	for _, cur_price := range stockPrice {
		if cur_profit := cur_price - purchase; cur_profit > profit {
			profit = cur_profit
		} else if purchase > cur_price {
			purchase = cur_price
		}
	}
	return profit
}
