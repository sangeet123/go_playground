package bubbleinttest

import (
	"go_playground/sort"
	"reflect"
	"testing"
)

func TestQuickSortInt(t *testing.T) {
	nos := []int{5, 4, 3, 2, 1}
	expected := []int{1, 2, 3, 4, 5}
	sort.QuickSortInt(nos)
	if !reflect.DeepEqual(expected, nos) {
		t.Error("Expected ", expected, " but received ", nos)
	}
}

func TestQuickSortIntEmptyArr(t *testing.T) {
	nos := []int{}
	expected := []int{}
	sort.QuickSortInt(nos)
	if !reflect.DeepEqual(expected, nos) {
		t.Error("Expected ", expected, " but received ", nos)
	}
}

func TestQuickSortIntAlreadySortedArr(t *testing.T) {
	nos := []int{1, 2, 3, 4, 5}
	expected := []int{1, 2, 3, 4, 5}
	sort.QuickSortInt(nos)
	if !reflect.DeepEqual(expected, nos) {
		t.Error("Expected ", expected, " but received ", nos)
	}
}
