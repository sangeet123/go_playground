package listprobtest

import (
	"go_playground/list"
	"go_playground/listprob"
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
	dupsdelete := listprob.RemoveDups(list)
	it := dupsdelete.Iterator()
	values := []int{1, 2, 3, 4}

	if dupsdelete.Size() != 4 {
		t.Error("Expected 4 but got ", list.Size())
	}

	for _, val := range values {
		val_in_list := it.Next()
		if val != val_in_list {
			t.Error("Expected ", val, " but received ", val_in_list)
		}
	}
}

func TestRemoveDupsNullList(t *testing.T) {
	list := new(list.List)
	dupsdelete := listprob.RemoveDups(list)

	if dupsdelete.Size() != 0 {
		t.Error("Expected 4 but got ", list.Size())
	}
}
