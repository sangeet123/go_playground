package sumpathtest

import (
	"go_playground/tree"
	"reflect"
	"testing"
)

func TestSumPathExample1(t *testing.T) {
	nos := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	tre := tree.CreateFullTree(nos)
	expected := [][]int{{9}, {3, 6}}

	result := tre.GetPath(9)
	if !reflect.DeepEqual(expected, result) {
		t.Error("expected inOrder ", expected, " but received ", result)
	}
}

func TestSumPathExample2(t *testing.T) {
	nos := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	tre := tree.CreateFullTree(nos)
	expected := [][]int{{1, 2, 4, 8}, {2, 4, 9}, {5, 10}}

	result := tre.GetPath(15)
	if !reflect.DeepEqual(expected, result) {
		t.Error("expected inOrder ", expected, " but received ", result)
	}
}

func TestSumPathExample3(t *testing.T) {
	nos := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	tre := tree.CreateFullTree(nos)
	expected := [][]int{{1}}

	result := tre.GetPath(1)
	if !reflect.DeepEqual(expected, result) {
		t.Error("expected inOrder ", expected, " but received ", result)
	}
}

func TestSumPathExample4(t *testing.T) {
	nos := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	tre := tree.CreateFullTree(nos)
	expected := [][]int{{1, 2}, {3}}

	result := tre.GetPath(3)
	if !reflect.DeepEqual(expected, result) {
		t.Error("expected inOrder ", expected, " but received ", result)
	}
}
