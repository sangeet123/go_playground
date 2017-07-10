package arrayprobtest

import (
	"go_playground/arrayprob"
	"testing"
)

func TestFourSumExample1(t *testing.T) {
	arr := []int{100, 4, 200, 1, 3, 2}
	if received := arrayprob.LengthOfLongestSequence(arr); received != 4 {
		t.Error("Expected 4 but received ", received)
	}
}

func TestFourSumExample2(t *testing.T) {
	arr := []int{100, 4, 200, 1, 3, 2, 101, 102, 104, 97, 10000, 2222, 98, 4444, 99, 8888}
	if received := arrayprob.LengthOfLongestSequence(arr); received != 6 {
		t.Error("Expected 4 but received ", received)
	}
}

func TestFourSumExample3(t *testing.T) {
	arr := []int{2, 4, 6, 8, 10}
	if received := arrayprob.LengthOfLongestSequence(arr); received != 1 {
		t.Error("Expected 4 but received ", received)
	}
}

func TestFourSumExample4(t *testing.T) {
	arr := []int{}
	if received := arrayprob.LengthOfLongestSequence(arr); received != 0 {
		t.Error("Expected 4 but received ", received)
	}
}
