package selectioninttest

import (
	"go_playground/sort"
	"reflect"
	"testing"
)

func TestInsertionSortInt(t *testing.T) {
	nos := []int{5, 4, 3, 2, 1}
	expected := []int{1, 2, 3, 4, 5}
	sort.InsertionSortInt(nos)
	if !reflect.DeepEqual(expected, nos) {
		t.Error("Expected ", expected, " but received ", nos)
	}
}

func TestInsertionSortIntEmptyArr(t *testing.T) {
	nos := []int{}
	expected := []int{}
	sort.InsertionSortInt(nos)
	if !reflect.DeepEqual(expected, nos) {
		t.Error("Expected ", expected, " but received ", nos)
	}
}
