package addtwolisttest

import (
	"go_playground/list"
	"testing"
)

func prepare_list() *list.List {
	list := new(list.List)
	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	list.Insert(4)
	return list
}

func TestReverse(t *testing.T) {
	l := prepare_list()
	l.Reverse()
	it := l.Iterator()
	values := []int{1, 2, 3, 4}

	if l.Size() != 4 {
		t.Error("Expected 4 but got ", l.Size())
	}

	for _, val := range values {
		val_in_list := it.Next()
		if val != val_in_list {
			t.Error("Expected ", val, " but received ", val_in_list)
		}
	}

	//verify the state from backwards
	bit := l.Iterator()
	for i := 0; i < l.Size()-1; i++ {
		bit.Next()
	}

	for i := l.Size() - 1; i >= 0; i-- {
		val_in_list := bit.Prev()
		if values[i] != val_in_list {
			t.Error("Expected ", values[i], " but received ", val_in_list)
		}
	}

}

func TestReverseForEmptyList(t *testing.T) {
	l := new(list.List)
	l.Reverse()

	if l.Size() != 0 {
		t.Error("Expected 4 but got ", l.Size())
	}
}

func TestReverseForSingledItemList(t *testing.T) {
	l := new(list.List)
	l.Append(3)
	l.Reverse()

	if l.Size() != 1 {
		t.Error("Expected 1 but got ", l.Size())
	}

	if l.Get(0) != 3 {
		t.Error("Expected 3 but got ", l.Get(0))
	}
}
