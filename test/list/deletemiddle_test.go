package kthtolasttest

import (
	"go_playground/list"
	"testing"
)

func prepare_list() *list.List {
	list := new(list.List)
	list.Append(1)
	list.Append(2)
	list.Append(3)
	return list
}

func TestEmptyList(t *testing.T) {
	list := new(list.List)
	list.DeleteMiddle()

	if list.Size() != 0 {
		t.Error("Expected 0 but got ", list.Size())
	}
}

func TestListWithSingleItem(t *testing.T) {
	list := new(list.List)
	list.Append(1)
	list.DeleteMiddle()

	if list.Size() != 0 {
		t.Error("Expected 0 but got ", list.Size())
	}
}

func TestListWithTwoItems(t *testing.T) {
	list := new(list.List)
	list.Append(1)
	list.Append(2)
	list.DeleteMiddle()

	if list.Size() != 1 {
		t.Error("Expected 1 but got ", list.Size())
	}

	if list.Get(0) != 2 {
		t.Error("Expected 2 but got ", list.Get(0))
	}
}

func TestDeleteMiddleOddSizeList(t *testing.T) {
	list := prepare_list()
	list.DeleteMiddle()
	it := list.Iterator()
	values := []int{1, 3}

	if list.Size() != 2 {
		t.Error("Expected 2 but got ", list.Size())
	}

	for _, val := range values {
		val_in_list := it.Next()
		if val != val_in_list {
			t.Error("Expected ", val, " but received ", val_in_list)
		}
	}
}

func TestDeleteMiddleEvenSizeList(t *testing.T) {
	list := prepare_list()
	list.Append(4)
	list.DeleteMiddle()
	it := list.Iterator()
	values := []int{1, 2, 4}

	if list.Size() != 3 {
		t.Error("Expected 3 but got ", list.Size())
	}

	for _, val := range values {
		val_in_list := it.Next()
		if val != val_in_list {
			t.Error("Expected ", val, " but received ", val_in_list)
		}
	}
}
