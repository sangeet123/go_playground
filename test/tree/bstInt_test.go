package bstinttest

import (
	"go_playground/tree"
	"reflect"
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

func TestCharArrayWithSpace(t *testing.T) {
	nos := []int{4, 5, 3, 2, 7, 8, 1, 2}
	expected := []int{1, 2, 2, 3, 4, 5, 7, 8}
	tre := tree.GetIntTree(comparator)
	for _, val := range nos {
		tre.Insert(val)
	}

	inorder := tre.InOrder()
	if !reflect.DeepEqual(expected, inorder) {
		t.Error("expected ", expected, " but received ", inorder)
	}

}
