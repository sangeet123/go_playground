package isbsttest

import (
	"go_playground/tree"
	"testing"
)

func TestIsBSTEmptyArray(t *testing.T) {
	nos := []int{}
	tre := tree.CreateFullTree(nos)

	if !tre.IsBST() {
		t.Error("expected true but received false")
	}
}

func TestIsBSTSingleElementTree(t *testing.T) {
	nos := []int{1}
	tre := tree.CreateFullTree(nos)

	if !tre.IsBST() {
		t.Error("expected true but received false")
	}
}

func TestIsNotBST(t *testing.T) {
	nos := []int{10, 5, 15, 3, 4}
	tre := tree.CreateFullTree(nos)

	if tre.IsBST() {
		t.Error("expected false but received true")
	}
}

func TestIsBST(t *testing.T) {
	nos := []int{10, 5, 15, 3, 7}
	tre := tree.CreateFullTree(nos)

	if !tre.IsBST() {
		t.Error("expected true but received false")
	}
}

func TestIsBSTWithDuplicatesOnLeftNode(t *testing.T) {
	nos := []int{10, 5, 15, 5, 7}
	tre := tree.CreateFullTree(nos)

	if !tre.IsBST() {
		t.Error("expected true but received false")
	}
}

func TestIsBSTWithDuplicatesOnRightNode(t *testing.T) {
	nos := []int{10, 5, 15, 5, 7, 12, 15}
	tre := tree.CreateFullTree(nos)

	if tre.IsBST() {
		t.Error("expected false but received true")
	}
}
