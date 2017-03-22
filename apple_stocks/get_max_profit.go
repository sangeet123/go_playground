package applestock

func get_max_profit(stocks_price []int)int{
	if len(stocks_price) == 0{
		return 0
	}

	profit := 0
	purchase := stocks_price[0]

	for _, cur_price := range stocks_price{
		profit_on_cur_transaction := cur_price - purchase
		if profit_on_cur_transaction > profit{
			profit = profit_on_cur_transaction
		}else if purchase > cur_price{
			purchase = cur_price
		}
	}
	return profit
}