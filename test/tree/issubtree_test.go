package issubtreetest

import (
	"go_playground/tree"
	"testing"
)

func TestIsSubTreeForNonEmptySameArray(t *testing.T) {
	nos := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	t1 := tree.CreateFullTree(nos)
	t2 := tree.CreateFullTree(nos)

	if !tree.IsSubTree(t1, t2) {
		t.Error("expected true but received false ")
	}
}

func TestIsSubTreeForEmptyArray(t *testing.T) {
	nos := []int{}
	t1 := tree.CreateFullTree(nos)
	t2 := tree.CreateFullTree(nos)

	if !tree.IsSubTree(t1, t2) {
		t.Error("expected true but received false ")
	}
}

func TestIsSubTreeForNonEmptyDifferentArrayExample1(t *testing.T) {
	nos1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	nos2 := []int{2, 4, 5, 8, 9, 10, 11}
	t1 := tree.CreateFullTree(nos1)
	t2 := tree.CreateFullTree(nos2)

	if !tree.IsSubTree(t1, t2) {
		t.Error("expected true but received false ")
	}
}

func TestIsSubTreeForNonEmptyDifferentArrayExample2(t *testing.T) {
	nos1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	nos2 := []int{9}
	t1 := tree.CreateFullTree(nos1)
	t2 := tree.CreateFullTree(nos2)

	if !tree.IsSubTree(t1, t2) {
		t.Error("expected true but received false ")
	}
}

func TestIsSubTreeForNonEmptyDifferentArrayExample3(t *testing.T) {
	nos1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	nos2 := []int{6, 12}
	t1 := tree.CreateFullTree(nos1)
	t2 := tree.CreateFullTree(nos2)

	if !tree.IsSubTree(t1, t2) {
		t.Error("expected true but received false ")
	}
}

func TestIsNotSubTreeForNonEmptyExample1(t *testing.T) {
	nos1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	nos2 := []int{2, 4, 5}
	t1 := tree.CreateFullTree(nos1)
	t2 := tree.CreateFullTree(nos2)

	if tree.IsSubTree(t1, t2) {
		t.Error("expected true but received false ")
	}
}

func TestIsNotSubTreeForNonEmptyExample2(t *testing.T) {
	nos1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	nos2 := []int{3, 6, 7}
	t1 := tree.CreateFullTree(nos1)
	t2 := tree.CreateFullTree(nos2)

	if tree.IsSubTree(t1, t2) {
		t.Error("expected true but received false ")
	}
}

func TestIsNotSubTreeForNonEmptyExample3(t *testing.T) {
	nos1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	nos2 := []int{1}
	t1 := tree.CreateFullTree(nos1)
	t2 := tree.CreateFullTree(nos2)

	if tree.IsSubTree(t1, t2) {
		t.Error("expected true but received false ")
	}
}
