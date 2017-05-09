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
	expectedInOrder := []int{1, 2, 2, 3, 4, 5, 7, 8}
	expectedPreOrder := []int{4, 3, 2, 1, 2, 5, 7, 8}
	expectedPostOrder := []int{2, 1, 2, 3, 8, 7, 5, 4}
	tre := tree.GetIntTree(comparator)
	for _, val := range nos {
		tre.Insert(val)
	}

	inorder := tre.InOrder()
	if !reflect.DeepEqual(expectedInOrder, inorder) {
		t.Error("expected inOrder ", expectedInOrder, " but received ", inorder)
	}

	preOrder := tre.PreOrder()
	if !reflect.DeepEqual(expectedPreOrder, preOrder) {
		t.Error("expected preOrder ", expectedPreOrder, " but received ", preOrder)
	}

	postOrder := tre.PostOrder()
	if !reflect.DeepEqual(expectedPostOrder, postOrder) {
		t.Error("expected postOrder ", expectedPostOrder, " but received ", postOrder)
	}

}
