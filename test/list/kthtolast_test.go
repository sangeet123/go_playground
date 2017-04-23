package kthtolasttest

import (
	"go_playground/list"
	"reflect"
	"testing"
)

func prepare_list() *list.List {
	list := new(list.List)
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	return list
}

func Test1thToLast(t *testing.T) {
	list := prepare_list()
	data := list.KthToLast(1)
	if !reflect.DeepEqual(3, data) {
		t.Error("Expected 3 but got ", data)
	}
}

func Test0thToLast(t *testing.T) {
	list := prepare_list()
	data := list.KthToLast(0)
	if !reflect.DeepEqual(4, data) {
		t.Error("Expected 3 but got ", data)
	}
}

func Test3thToLast(t *testing.T) {
	list := prepare_list()
	data := list.KthToLast(3)
	if !reflect.DeepEqual(1, data) {
		t.Error("Expected 3 but got ", data)
	}
}

func Test5thToLast(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic but failed")
		}
	}()

	list := prepare_list()
	list.KthToLast(5)
}

func TestMinus1thToLast(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic but failed")
		}
	}()

	list := prepare_list()
	list.KthToLast(5)
}
