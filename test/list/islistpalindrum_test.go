package islistpalindrum

import (
	"go_playground/list"
	"testing"
)

func TestEmptyList(t *testing.T) {
	list := new(list.List)
	expected := list.IsListPalindrum()

	if expected != false {
		t.Error("Expected false but got ", expected)
	}
}

func TestListWithOneItem(t *testing.T) {
	list := new(list.List)
	list.Append(1)
	expected := list.IsListPalindrum()

	if expected != true {
		t.Error("Expected true but got ", expected)
	}
}

func TestPalindrumListWithEvenItems(t *testing.T) {
	list := new(list.List)
	list.Append(1)
	list.Append(1)
	expected := list.IsListPalindrum()

	if expected != true {
		t.Error("Expected true but got ", expected)
	}
}

func TestNotPalindrumListWithEvenItems(t *testing.T) {
	list := new(list.List)
	list.Append(1)
	list.Append(1)
	list.Append(2)
	list.Append(1)
	expected := list.IsListPalindrum()

	if expected != false {
		t.Error("Expected false but got ", expected)
	}
}

func TestPalindrumListWithOddItems(t *testing.T) {
	list := new(list.List)
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(2)
	list.Append(1)
	expected := list.IsListPalindrum()

	if expected != true {
		t.Error("Expected true but got ", expected)
	}
}

func TestNotPalindrumListWithOddItems(t *testing.T) {
	list := new(list.List)
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(3)
	list.Append(1)
	expected := list.IsListPalindrum()

	if expected != false {
		t.Error("Expected false but got ", expected)
	}
}
