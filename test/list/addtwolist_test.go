package addtwolisttest

import (
	"go_playground/list"
	"testing"
)

func TestAddEmptyLists(t *testing.T) {
	l1 := new(list.List)
	l2 := new(list.List)
	l3 := list.AddList(l1, l2)

	if l3.Size() != 0 {
		t.Error("Expected 0 but got ", l3.Size())
	}
}

func TestAddListsOfEqualSize(t *testing.T) {
	l1 := new(list.List)
	// 654
	l1.Insert(4)
	l1.Insert(5)
	l1.Insert(6)

	// 387
	l2 := new(list.List)
	l2.Insert(7)
	l2.Insert(8)
	l2.Insert(3)

	l3 := list.AddList(l1, l2)

	if l3.Size() != 4 {
		t.Error("Expected 4 but got ", l3.Size())
	}

	values := []int{1, 0, 4, 1}
	it := l3.Iterator()
	for _, val := range values {
		val_in_list := it.Next()
		if val != val_in_list {
			t.Error("Expected ", val, " but received ", val_in_list)
		}
	}
}

func TestAddListsOfUnEqualSize(t *testing.T) {
	l1 := new(list.List)
	// 999
	l1.Insert(9)
	l1.Insert(9)
	l1.Insert(9)

	// 12
	l2 := new(list.List)
	l2.Insert(2)
	l2.Insert(1)

	l3 := list.AddList(l1, l2)

	if l3.Size() != 4 {
		t.Error("Expected 4 but got ", l3.Size())
	}

	values := []int{1, 0, 1, 1}
	it := l3.Iterator()
	for _, val := range values {
		val_in_list := it.Next()
		if val != val_in_list {
			t.Error("Expected ", val, " but received ", val_in_list)
		}
	}
}
