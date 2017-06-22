package isBalancedtest

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

func TestEmptyTree(t *testing.T) {
	tre := tree.GetIntTree(comparator)
	if !tre.IsBalanced() {
		t.Error("Expected true but received false")
	}
}

func TestTreeWithSingleNode(t *testing.T) {
	tre := tree.GetIntTree(comparator)
	tre.Insert(2)
	if !tre.IsBalanced() {
		t.Error("Expected true but received false")
	}
}

func TestNotBalancedTree(t *testing.T) {
	nos := []int{56, 47, 68, 45, 50, 48, 52}
	tre := tree.GetIntTree(comparator)
	for _, val := range nos {
		tre.Insert(val)
	}
	if tre.IsBalanced() {
		t.Error("Expected false but received true")
	}
}

func TestBalancedTree(t *testing.T) {
	nos := []int{56, 47, 68, 45, 50, 65, 70}
	tre := tree.GetIntTree(comparator)
	for _, val := range nos {
		tre.Insert(val)
	}
	if !tre.IsBalanced() {
		t.Error("Expected true but received false")
	}
}

func TestSkewedRightTree(t *testing.T) {
	nos := []int{1, 2, 3, 4, 5, 6, 7, 8}
	tre := tree.GetIntTree(comparator)
	for _, val := range nos {
		tre.Insert(val)
	}
	if tre.IsBalanced() {
		t.Error("Expected false but received true")
	}
}

func TestSkewedLefttTree(t *testing.T) {
	nos := []int{8, 7, 6, 5, 4, 3, 2, 1}
	tre := tree.GetIntTree(comparator)
	for _, val := range nos {
		tre.Insert(val)
	}
	if tre.IsBalanced() {
		t.Error("Expected false but received true")
	}
}
