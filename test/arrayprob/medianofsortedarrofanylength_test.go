package medianofsortedarrtest

import (
	"go_playground/arrayprob"
	"testing"
)

func TestArrWithTwoElements(t *testing.T) {
	arr1 := []int{1, 2}
	arr2 := []int{3, 4}
	if val := arrayprob.GetMedianOfUnequalLengthArr(arr1, arr2); val != 2.5 {
		t.Error("Expected 2.5 but received ", val)
	}
}

func TestArrWithSixElements(t *testing.T) {
	arr1 := []int{1, 2, 3, 4, 7, 10}
	arr2 := []int{5, 7, 9, 10, 11, 12}
	if val := arrayprob.GetMedianOfUnequalLengthArr(arr1, arr2); val != 7.0 {
		t.Error("Expected 2.5 but received ", val)
	}
}

func TestArrWithUnEqualElementsZumbleup(t *testing.T) {
	arr1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 13, 14, 15, 20, 21, 22}
	arr2 := []int{10, 11, 16, 17, 18, 19}
	if val := arrayprob.GetMedianOfUnequalLengthArr(arr1, arr2); val != 11.5 {
		t.Error("Expected 11.5 but received ", val)
	}
}

func TestArrWithOneMedianElementsZumbleup(t *testing.T) {
	arr1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 13, 14, 15, 20, 21, 22, 23}
	arr2 := []int{10, 11, 16, 17, 18, 19}
	if val := arrayprob.GetMedianOfUnequalLengthArr(arr1, arr2); val != 12 {
		t.Error("Expected 12 but received ", val)
	}
}
