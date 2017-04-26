package checkintersectedlisttest

import (
	"go_playground/list"
	"testing"
)

func TestIntersectedList(t *testing.T) {
	l1, l2, pos1, pos2 := list.GetIntersectedList()

	a := l1.Get(pos1 - 1)
	b := l2.Get(pos2 - 1)

	if a != b {
		t.Error("Expected equal but got ", a, b)
	}

	intersected, val := list.AreListIntersected(l1, l2)

	if intersected != true && val != a {
		t.Error("Expected true but got false with expected value  but got ", a, val)
	}
}

func TestEmptyListThatAreNotInterSected(t *testing.T) {
	l1 := new(list.List)
	l2 := new(list.List)

	intersected, _ := list.AreListIntersected(l1, l2)

	if intersected != false {
		t.Error("Expected false but got true")
	}

}

func TestListThatAreNotInterSected(t *testing.T) {
	l1 := new(list.List)
	l1.Append(1)
	l1.Append(2)
	l1.Append(3)
	l2 := new(list.List)
	l2.Append(1)
	l2.Append(2)

	intersected, _ := list.AreListIntersected(l1, l2)

	if intersected != false {
		t.Error("Expected false but got true")
	}

}
