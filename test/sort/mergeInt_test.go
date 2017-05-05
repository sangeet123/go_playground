package bubbleinttest

import (
	"go_playground/sort"
	"reflect"
	"testing"
)

func TestMergeSortInt(t *testing.T) {
	nos := []int{5, 4, 3, 2, 1}
	expected := []int{1, 2, 3, 4, 5}
	sort.MergeSortInt(nos)
	if !reflect.DeepEqual(expected, nos) {
		t.Error("Expected ", expected, " but received ", nos)
	}
}

func TestMergeSortIntEmptyArr(t *testing.T) {
	nos := []int{}
	expected := []int{}
	sort.MergeSortInt(nos)
	if !reflect.DeepEqual(expected, nos) {
		t.Error("Expected ", expected, " but received ", nos)
	}
}

func TestMergeSortIntArrWithOneElement(t *testing.T) {
	nos := []int{1}
	expected := []int{1}
	sort.MergeSortInt(nos)
	if !reflect.DeepEqual(expected, nos) {
		t.Error("Expected ", expected, " but received ", nos)
	}
}

func TestMergeSortIntArrWithTwoElements(t *testing.T) {
	nos := []int{2, 1}
	expected := []int{1, 2}
	sort.MergeSortInt(nos)
	if !reflect.DeepEqual(expected, nos) {
		t.Error("Expected ", expected, " but received ", nos)
	}
}
