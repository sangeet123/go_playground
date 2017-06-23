package waystoconstructbsttest

import (
	"go_playground/tree"
	"testing"
)

func TestWaysToConstructBSTEmptyTree(t *testing.T) {
	nos := []int{}

	if ways := tree.WaysToConstructBst(nos); ways != 1 {
		t.Error("expected 1 but got ", ways)
	}
}

func TestWaysToConstructBSTSingleElementTree(t *testing.T) {
	nos := []int{1}

	if ways := tree.WaysToConstructBst(nos); ways != 1 {
		t.Error("expected 1 but got ", ways)
	}
}

func TestWaysToConstructBSTTwoElementsTree(t *testing.T) {
	nos := []int{1, 2}

	if ways := tree.WaysToConstructBst(nos); ways != 1 {
		t.Error("expected 1 but got ", ways)
	}
}

func TestWaysToConstructBSTThreeElementsTree(t *testing.T) {
	nos := []int{2, 1, 3}

	if ways := tree.WaysToConstructBst(nos); ways != 2 {
		t.Error("expected 2 but got ", ways)
	}
}

func TestWaysToConstructBSTExample1(t *testing.T) {
	nos := []int{4, 2, 8, 1, 3}

	if ways := tree.WaysToConstructBst(nos); ways != 8 {
		t.Error("expected 4 but got ", ways)
	}
}

func TestWaysToConstructBSTExample2(t *testing.T) {
	nos := []int{4, 2, 6, 1, 3, 5, 7}

	if ways := tree.WaysToConstructBst(nos); ways != 80 {
		t.Error("expected 80 but got ", ways)
	}
}
