package splitaroundktest

import (
	"go_playground/list"
	"testing"
)

func prepare_list() *list.List {
	list := new(list.List)
	list.Append(4)
	list.Append(5)
	list.Append(6)
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(2)
	return list
}

func TestSplitAroundK(t *testing.T) {
	list := prepare_list()
	list.SplitAroundK(6)
	it := list.Iterator()
	values := []int{4,5,1,2,3,2,6}

	if list.Size() != 7 {
		t.Error("Expected 7 but got ", list.Size())
	}

	for _, val := range values {
		val_in_list := it.Next()
		if val != val_in_list {
			t.Error("Expected ", val, " but received ", val_in_list)
		}
	}
}

func TestSplitAroundKEmptyList(t *testing.T) {
	list := new(list.List)
	list.SplitAroundK(1)

	if list.Size() != 0 {
		t.Error("Expected 4 but got ", list.Size())
	}
}

func TestSplitAroundKSingleItemList(t *testing.T) {
	list := new(list.List)
	list.Append(1)
	list.SplitAroundK(1)

	if list.Size() != 1 {
		t.Error("Expected 4 but got ", list.Size())
	}

	if list.Get(0) != 1{
		t.Error("Expected 1 but got ", list.Get(0))
	}
}


