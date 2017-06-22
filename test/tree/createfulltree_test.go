package bstinttest

import (
	"go_playground/tree"
	"reflect"
	"testing"
)

func TestCreateFullTree(t *testing.T) {
	nos := []int{4, 5, 3, 2, 7, 8, 1, 2}
	tre := tree.CreateFullTree(nos)

	levelOrder := tre.LevelOrder()
	if !reflect.DeepEqual(nos, levelOrder) {
		t.Error("expected levelorder ", nos, " but received ", levelOrder)
	}
}
