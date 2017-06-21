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

	if _, ok := tre.InOrderSuccessor(1); ok {
		t.Error("Expected false but received true")
	}
}

func TestTreeWithSingleNode(t *testing.T) {
	tre := tree.GetIntTree(comparator)
	tre.Insert(1)
	if _, ok := tre.InOrderSuccessor(1); ok {
		t.Error("Expected false but received true")
	}
}

func TestTreeWithRootAndLeftChild(t *testing.T) {
	tre := tree.GetIntTree(comparator)
	tre.Insert(2)
	tre.Insert(1)
	if val, ok := tre.InOrderSuccessor(1); !ok || val != 2 {
		t.Error("Expected true, 2  but received false and ", val)
	}
}

func TestTreeWithRootAndRightChild(t *testing.T) {
	tre := tree.GetIntTree(comparator)
	tre.Insert(1)
	tre.Insert(2)
	if val, ok := tre.InOrderSuccessor(1); !ok || val != 2 {
		t.Error("Expected true, 2  but received false and ", val)
	}
}

func TestTreeWithRightChild(t *testing.T) {
	tre := tree.GetIntTree(comparator)
	nos := []int{56, 37, 24, 85, 11, 10, 9, 3, 2, 1, 30, 29, 28}
	for _, val := range nos {
		tre.Insert(val)
	}
	if val, ok := tre.InOrderSuccessor(24); !ok || val != 28 {
		t.Error("Expected false but received true")
	}
}

func TestRighSkewedTree(t *testing.T) {
	tre := tree.GetIntTree(comparator)
	nos := []int{56, 57, 58, 59, 60}
	for _, val := range nos {
		tre.Insert(val)
	}
	for i := 0; i < len(nos)-1; i++ {
		if in, ok := tre.InOrderSuccessor(nos[i]); !ok || in != nos[i+1] {
			t.Error("Expected true,", nos[i+1], " but received false and ", in)
		}
	}
	if val, ok := tre.InOrderSuccessor(60); ok {
		t.Error("Expected false but received true and ", val)
	}

}

func TestTreeithLeftHavingRightChild(t *testing.T) {
	tre := tree.GetIntTree(comparator)
	nos := []int{56, 37, 24, 85, 11, 10, 9, 3, 2, 1, 30, 29, 28, 60, 65, 66, 67, 68, 69}
	for _, val := range nos {
		tre.Insert(val)
	}
	if val, ok := tre.InOrderSuccessor(69); !ok || val != 85 {
		t.Error("Expected true, 85  but received false and ", val)
	}
}
