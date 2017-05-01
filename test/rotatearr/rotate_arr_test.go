package rotatedarrtest

import (
	"go_playground/rotatearr"
	"reflect"
	"testing"
)

func TestRotateArrayByFactorWithinLength(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	expected := []int{5, 1, 2, 3, 4}
	rotatearr.Rotate(arr, 1)
	if !reflect.DeepEqual(expected, arr) {
		t.Error("Expected ", expected, " but got ", arr)
	}
}

func TestRotateArrayByFactorExceedingLength(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	expected := []int{4, 5, 1, 2, 3}
	rotatearr.Rotate(arr, 7)
	if !reflect.DeepEqual(expected, arr) {
		t.Error("Expected ", expected, " but got ", arr)
	}
}

func TestRotateArrayThatDoesNotChange(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	expected := []int{1, 2, 3, 4, 5}
	rotatearr.Rotate(arr, 10)
	if !reflect.DeepEqual(expected, arr) {
		t.Error("Expected ", expected, " but got ", arr)
	}
}

func TestRotateArrayByFactorNegativeVal(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected error but no error ocurred")
		}
	}()
	rotatearr.Rotate(arr, -1)
}
