package arrayprobtest

import (
	"go_playground/arrayprob"
	"reflect"
	"testing"
)

func TestRemoveDupsFromSortedZeroDistApart(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	received := arrayprob.RemoveDups(arr, 0)
	expected := []int(nil)
	if !reflect.DeepEqual(expected, received) {
		t.Error("Expected ", expected, " but received ", received)
	}
}

func TestRemoveDupsFromSortedNilArray(t *testing.T) {
	received := arrayprob.RemoveDups(nil, 5)
	expected := []int(nil)
	if !reflect.DeepEqual(expected, received) {
		t.Error("Expected ", expected, " but received ", received)
	}
}

func TestRemoveDupsFromSortedArrayWithDuplicatesExample1(t *testing.T) {
	received := arrayprob.RemoveDups([]int{1, 1, 1, 2, 2, 3, 3}, 1)
	expected := []int{1, 2, 3}
	if !reflect.DeepEqual(expected, received) {
		t.Error("Expected ", expected, " but received ", received)
	}
}

func TestRemoveDupsFromSortedArrayWithDuplicatesExample2(t *testing.T) {
	received := arrayprob.RemoveDups([]int{1, 1, 1, 2, 2, 3, 3, 3}, 2)
	expected := []int{1, 1, 2, 2, 3, 3}
	if !reflect.DeepEqual(expected, received) {
		t.Error("Expected ", expected, " but received ", received)
	}
}

func TestRemoveDupsFromSortedArrayWithDuplicatesExample3(t *testing.T) {
	received := arrayprob.RemoveDups([]int{1, 1, 1, 2, 2, 2, 2, 3, 3, 3}, 3)
	expected := []int{1, 1, 1, 2, 2, 2, 3, 3, 3}
	if !reflect.DeepEqual(expected, received) {
		t.Error("Expected ", expected, " but received ", received)
	}
}
