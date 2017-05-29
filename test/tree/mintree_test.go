package mintreetest

import (
	"go_playground/tree"
	"reflect"
	"testing"
)

func TestTreeInsertOperation(t *testing.T) {
	nos := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	tre := tree.CreateMinimumTree(nos)
	inorder := tre.InOrder()
	if !reflect.DeepEqual(nos, inorder) {
		t.Error("expected inOrder ", nos, " but received ", inorder)
	}

	height := tre.GetHeight()
	if height != 4 {
		t.Error("expected 3 but received ", height)
	}
}
