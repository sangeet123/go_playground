package arrayprobtest

import (
	"go_playground/arrayprob"
	"reflect"
	"testing"
)

func TestFourSumExample1(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	expected := []int{0, 1, 2, 3}
	if hasFourSum, received := arrayprob.FourSumArr(arr, 10); !hasFourSum || !reflect.DeepEqual(expected, received) {
		t.Error("Expected true and ", expected, " but received false and ", received)
	}
}

func TestFourSumExample2(t *testing.T) {
	arr := []int{1, 2, 3}
	expected := []int{-1, -1, -1, -1}
	if hasFourSum, received := arrayprob.FourSumArr(arr, 10); hasFourSum || !reflect.DeepEqual(expected, received) {
		t.Error("Expected true and ", expected, " but received false and ", received)
	}
}

func TestFourSumExample3(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 20}
	expected := []int{0, 1, 3, 8}
	if hasFourSum, received := arrayprob.FourSumArr(arr, 16); !hasFourSum || !reflect.DeepEqual(expected, received) {
		t.Error("Expected true and ", expected, " but received false and ", received)
	}
}
