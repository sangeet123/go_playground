package bstinttest

import (
	"go_playground/tree"
	"reflect"
	"testing"
)

func TestInvertTreeExp1(t *testing.T) {
	nos := []int{4, 5, 3, 2, 7, 8, 1, 2}
	expected := []int{4, 3, 5, 1, 8, 7, 2, 2}
	tre := tree.CreateFullTree(nos)
	tre.Invert()
	levelOrder := tre.LevelOrder()
	if !reflect.DeepEqual(expected, levelOrder) {
		t.Error("expected levelorder ", expected, " but received ", levelOrder)
	}
}

func TestInvertTreeExp2(t *testing.T) {
	nos := []int{4}
	expected := []int{4}
	tre := tree.CreateFullTree(nos)
	tre.Invert()
	levelOrder := tre.LevelOrder()
	if !reflect.DeepEqual(expected, levelOrder) {
		t.Error("expected levelorder ", expected, " but received ", levelOrder)
	}
}

func TestInvertTreeExp3(t *testing.T) {
	nos := []int{}
	expected := []int{}
	tre := tree.CreateFullTree(nos)
	tre.Invert()
	levelOrder := tre.LevelOrder()
	if !reflect.DeepEqual(expected, levelOrder) {
		t.Error("expected levelorder ", expected, " but received ", levelOrder)
	}
}
