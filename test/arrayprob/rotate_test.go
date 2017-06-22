package arrayprobtest

import (
	"go_playground/arrayprob"
	"reflect"
	"testing"
)

func TestRotateBy(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	arrayprob.RotateRightBy(arr, 3)
	expected := []int{3, 4, 5, 1, 2}
	if !reflect.DeepEqual(expected, arr) {
		t.Error("Expected ", expected, " but received ", arr)
	}
}

func TestRotateEmptyArray(t *testing.T) {
	arr := []int{}
	arrayprob.RotateRightBy(arr, 3)
	expected := []int{}
	if !reflect.DeepEqual(expected, arr) {
		t.Error("Expected ", expected, " but received ", arr)
	}
}

func TestRotateEmptyArrayByZero(t *testing.T) {
	arr := []int{}
	arrayprob.RotateRightBy(arr, 0)
	expected := []int{}
	if !reflect.DeepEqual(expected, arr) {
		t.Error("Expected ", expected, " but received ", arr)
	}
}

func TestRotateNonEmptyArrayByZero(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	arrayprob.RotateRightBy(arr, 0)
	expected := []int{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(expected, arr) {
		t.Error("Expected ", expected, " but received ", arr)
	}
}

func TestRotateNonEmptyArrayBySizeGreaterThanArrayLength(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	arrayprob.RotateRightBy(arr, 7)
	expected := []int{4, 5, 1, 2, 3}
	if !reflect.DeepEqual(expected, arr) {
		t.Error("Expected ", expected, " but received ", arr)
	}
}

func TestRotateNonEmptyArrayByNegativeLength(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected error but no error ocurred")
		}
	}()
	arrayprob.RotateRightBy(arr, -7)
	expected := []int{4, 5, 1, 2, 3}
	if !reflect.DeepEqual(expected, arr) {
		t.Error("Expected ", expected, " but received ", arr)
	}
}
