package applestock

import "testing"

func TestStockForIncreasingArray(t *testing.T) {
	stock_prices := []int{4, 5, 6, 12, 18, 25}
	expected := 21
	received := get_max_profit(stock_prices)
	if expected != received {
		t.Error("Expected ", expected, " but received ", received)
	}
}

func TestStockForDecreasingArray(t *testing.T) {
	stock_prices := []int{25, 18, 12, 6, 5, 4}
	expected := 0
	received := get_max_profit(stock_prices)
	if expected != received {
		t.Error("Expected ", expected, " but received ", received)
	}
}

func TestStockForIncreasingAndDecreasingArrayExample1(t *testing.T) {
	stock_prices := []int{5, 6, 12, 14, 10, 4, 12}
	expected := 9
	received := get_max_profit(stock_prices)
	if expected != received {
		t.Error("Expected ", expected, " but received ", received)
	}
}

func TestStockForIncreasingAndDecreasingArrayExample2(t *testing.T) {
	stock_prices := []int{5, 6, 12, 14, 10, 4, 12, 16}
	expected := 12
	received := get_max_profit(stock_prices)
	if expected != received {
		t.Error("Expected ", expected, " but received ", received)
	}
}

func TestStockForIncreasingAndDecreasingArrayExample3(t *testing.T) {
	stock_prices := []int{5, 6, 12, 14, 10, 4, 12, 16, 20, 1, 2, 3, 4, 21}
	expected := 20
	received := get_max_profit(stock_prices)
	if expected != received {
		t.Error("Expected ", expected, " but received ", received)
	}
}
