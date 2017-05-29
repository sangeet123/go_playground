package treesorttest

import (
	"go_playground/sort"
	"reflect"
	"testing"
)

func TestBubbleSortInt(t *testing.T) {
	nos := []int{5, 4, 3, 2, 1}
	expected := []int{1, 2, 3, 4, 5}
	sort.TreeSortInt(nos)
	if !reflect.DeepEqual(expected, nos) {
		t.Error("Expected ", expected, " but received ", nos)
	}
}

func TestBubbleSortIntEmptyArr(t *testing.T) {
	nos := []int{}
	expected := []int{}
	sort.TreeSortInt(nos)
	if !reflect.DeepEqual(expected, nos) {
		t.Error("Expected ", expected, " but received ", nos)
	}
}
