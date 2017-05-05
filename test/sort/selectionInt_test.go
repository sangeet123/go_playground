package selectioninttest

import (
	"go_playground/sort"
	"reflect"
	"testing"
)

func TestSelectionSortInt(t *testing.T) {
	nos := []int{5, 4, 3, 2, 1}
	expected := []int{1, 2, 3, 4, 5}
	sort.SelectionSortInt(nos)
	if !reflect.DeepEqual(expected, nos) {
		t.Error("Expected ", expected, " but received ", nos)
	}
}

func TestSelectionSortIntEmptyArr(t *testing.T) {
	nos := []int{}
	expected := []int{}
	sort.SelectionSortInt(nos)
	if !reflect.DeepEqual(expected, nos) {
		t.Error("Expected ", expected, " but received ", nos)
	}
}
