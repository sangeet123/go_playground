package hascycletest

import (
	"go_playground/list"
	"testing"
)

func TestIntersectedList(t *testing.T) {
	lst, val := list.CreateListWithCycle()

	hasCycle, atVal := lst.HasCycle()

	if !hasCycle && val != atVal {
		t.Error("Expected true but got false and the cycle is at point ", val)
	}
}

func TestListThatHasNoCycle(t *testing.T) {
	lst := new(list.List)
	lst.Append(1)
	lst.Append(2)
	lst.Append(3)

	if hasCycle, _ := lst.HasCycle(); !hasCycle {
		t.Error("Expected false but got true")
	}

}
