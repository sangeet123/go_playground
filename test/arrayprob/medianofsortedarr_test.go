package medianofsortedarrtest

import (
	"go_playground/arrayprob"
	"testing"
)

func TestArrWithTwoElements(t *testing.T) {
	arr1 := []int{1, 2}
	arr2 := []int{3, 4}
	if val := arrayprob.GetMedian(arr1, arr2); val != 2 {
		t.Error("Expected 2 but received ", val)
	}
}

func TestArrWithTwoElementsButMixedSorted(t *testing.T) {
	arr1 := []int{1, 3}
	arr2 := []int{2, 4}
	if val := arrayprob.GetMedian(arr1, arr2); val != 2 {
		t.Error("Expected 2 but received ", val)
	}
}

func TestArrWithThreeElements(t *testing.T) {
	arr1 := []int{1, 5, 7}
	arr2 := []int{2, 4, 7}
	if val := arrayprob.GetMedian(arr1, arr2); val != 4 {
		t.Error("Expected 4 but received ", val)
	}
}

func TestArrWithFourElements(t *testing.T) {
	arr1 := []int{1, 5, 7, 10}
	arr2 := []int{2, 4, 9, 13}
	if val := arrayprob.GetMedian(arr1, arr2); val != 5 {
		t.Error("Expected 1 but received ", val)
	}
}

func TestArrWithNineElements(t *testing.T) {
	arr1 := []int{1, 5, 7, 10, 12, 15, 17, 18, 20}
	arr2 := []int{2, 4, 9, 11, 14, 14, 15, 16, 17}
	if val := arrayprob.GetMedian(arr1, arr2); val != 12 {
		t.Error("Expected 12 but received ", val)
	}
}

func TestArrWithTenElementsWithAllCaseSatisfying(t *testing.T) {
	arr1 := []int{1, 5, 7, 10, 12, 15, 17, 18, 20, 25}
	arr2 := []int{2, 4, 11, 11, 18, 19, 20, 21, 22, 24}
	if val := arrayprob.GetMedian(arr1, arr2); val != 15 {
		t.Error("Expected 15 but received ", val)
	}
}
