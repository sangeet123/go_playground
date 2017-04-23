package removedupstest

import (
	"go_playground/list"
	"testing"
)

func prepare_list() *list.List {
	list := new(list.List)
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	return list
}

func TestRemoveDups(t *testing.T) {
	list := prepare_list()
	list.RemoveDupsInplace()
	it := list.Iterator()
	values := []int{1, 2, 3, 4}

	if list.Size() != 4 {
		t.Error("Expected 4 but got ", list.Size())
	}

	for _, val := range values {
		val_in_list := it.Next()
		if val != val_in_list {
			t.Error("Expected ", val, " but received ", val_in_list)
		}
	}
}

func TestRemoveDupsEmptyList(t *testing.T) {
	list := new(list.List)
	list.RemoveDupsInplace()

	if list.Size() != 0 {
		t.Error("Expected 4 but got ", list.Size())
	}
}
