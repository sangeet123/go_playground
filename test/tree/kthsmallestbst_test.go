package kthsmallestelementtest

import (
	"go_playground/tree"
	"testing"
)

func comparator(a, b int) int {
	if a == b {
		return 0
	} else if a > b {
		return 1
	}
	return -1
}

func TestKthSmallestElementExample1(t *testing.T) {
	tre := tree.GetIntTree(comparator)
	tre.Insert(1)
	if val := tree.GetSmallest(tre, 1); val != 1 {
		t.Error("expected 1 but received value ", val)
	}
}

func TestKthSmallestElementExample2(t *testing.T) {
	tre := tree.GetIntTree(comparator)
	nos := []int{4, 5, 3, 2, 7, 8, 1, 2}
	inOrder := []int{1, 2, 3, 4, 5, 7, 8}
	for _, val := range nos {
		tre.Insert(val)
	}

	for i := 1; i <= len(inOrder); i++ {
		if val := tree.GetSmallest(tre, i); val != inOrder[i-1] {
			t.Error("expected ", inOrder[i-1], " but received value ", val)
		}
	}
}

func TestKthSmallestElementExample3(t *testing.T) {
	tre := tree.GetIntTree(comparator)
	tre.Insert(1)
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected error but no error ocurred")
		}
	}()
	if val := tree.GetSmallest(tre, 0); val != 1 {
		t.Error("expected 1 but received value ", val)
	}
}

func TestKthSmallestElementExample4(t *testing.T) {
	tre := tree.GetIntTree(comparator)
	tre.Insert(1)
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected error but no error ocurred")
		}
	}()
	if val := tree.GetSmallest(tre, 2); val != 1 {
		t.Error("expected 1 but received value ", val)
	}
}
