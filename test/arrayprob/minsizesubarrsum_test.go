package arrayprobtest

import (
	"go_playground/arrayprob"
	"reflect"
	"testing"
)

func TestExample1(t *testing.T) {
	arr := []int{2, 3, 1, 2, 4, 3}
	expected := []int{4, 3}
	if received := arrayprob.GetMinSizeSubArr(arr, 7); !reflect.DeepEqual(expected, received) {
		t.Error("Expected true and ", expected, " but received false and ", received)
	}
}

func TestExample2(t *testing.T) {
	arr := []int{10, 11, 12, 13, 14, 20}
	expected := []int{20}
	if received := arrayprob.GetMinSizeSubArr(arr, 19); !reflect.DeepEqual(expected, received) {
		t.Error("Expected true and ", expected, " but received false and ", received)
	}
}

func TestExample3(t *testing.T) {
	arr := []int{20, 11, 12, 13, 14, 18}
	expected := []int{20}
	if received := arrayprob.GetMinSizeSubArr(arr, 19); !reflect.DeepEqual(expected, received) {
		t.Error("Expected true and ", expected, " but received false and ", received)
	}
}

func TestExample4(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	expected := []int{1, 2, 3, 4, 5, 6, 7}
	if received := arrayprob.GetMinSizeSubArr(arr, 28); !reflect.DeepEqual(expected, received) {
		t.Error("Expected true and ", expected, " but received false and ", received)
	}
}

func TestExample5(t *testing.T) {
	arr := []int{1, 2, 3, 1, 1, 1, 1}
	expected := []int{1, 2, 3}
	if received := arrayprob.GetMinSizeSubArr(arr, 6); !reflect.DeepEqual(expected, received) {
		t.Error("Expected true and ", expected, " but received false and ", received)
	}
}

func TestExample6(t *testing.T) {
	arr := []int{}
	expected := []int{}
	if received := arrayprob.GetMinSizeSubArr(arr, 6); !reflect.DeepEqual(expected, received) {
		t.Error("Expected true and ", expected, " but received false and ", received)
	}
}

func TestExample7(t *testing.T) {
	arr := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	expected := []int{}
	if received := arrayprob.GetMinSizeSubArr(arr, 26); !reflect.DeepEqual(expected, received) {
		t.Error("Expected true and ", expected, " but received false and ", received)
	}
}
