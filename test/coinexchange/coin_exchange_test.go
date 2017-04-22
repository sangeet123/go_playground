package coinexchangetest

import (
	"go_playground/coinexchange"
	"testing"
)

func TestGetTotalWaysHavingSolutionWithCoinsNotsorted(t *testing.T) {
	coins := []int{3, 2, 1}
	total := 4
	expected := 4
	received := coinexchange.GetTotalWays(total, coins)

	if expected != received {
		t.Error("Expected ", expected, " but received ", received)
	}

}

func TestGetTotalWaysHavingNoSolutionWithCoinsNotsorted(t *testing.T) {
	coins := []int{12, 18}
	total := 4
	expected := 0
	received := coinexchange.GetTotalWays(total, coins)

	if expected != received {
		t.Error("Expected ", expected, " but received ", received)
	}

}

func TestGetTotalWaysHavingSolutionWithCoinssorted(t *testing.T) {
	coins := []int{12, 18}
	total := 36
	expected := 2
	received := coinexchange.GetTotalWays(total, coins)

	if expected != received {
		t.Error("Expected ", expected, " but received ", received)
	}

}

func TestGetTotalWaysHavingSolutionWithCoinsNotsortedExample1(t *testing.T) {
	//solutions
	// {1,1,1,1,1,1,1,1,1}, {1,1,1,1,1,1,1,1,2}, {1,1,1,1,1,2,2}, {1,1,1,1,2,2,2}
	// {1,1,2,2,2,2}, {2,2,2,2,2}, {1,3,2,2,2}, {1,1,1,3,2,2}, {1,1,1,1,1,3,2},{1,1,1,1,1,1,1,3}
	// {3,3,2,2}, {3,3,3,1}, {3,3,1,1,1,1}, {3,1,1,1,1,1,1,1}, {3,3,2,1,1}

	coins := []int{3, 2, 1}
	total := 10
	expected := 14
	received := coinexchange.GetTotalWays(total, coins)

	if expected != received {
		t.Error("Expected ", expected, " but received ", received)
	}

}
