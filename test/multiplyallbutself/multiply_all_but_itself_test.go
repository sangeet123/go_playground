package multiplyallbutselftest

import (
	"go_playground/multiplyallbutself"
	"reflect"
	"testing"
)

func TestMultiplyAllButItself(t *testing.T) {
	values := multiplyallbutself.Integers{1, 2, 3, 4, 5, 6, 7}
	expected := multiplyallbutself.Integers{5040, 2520, 1680, 1260, 1008, 840, 720}

	received := multiplyallbutself.MultiplyAllButSelf(values)

	if !reflect.DeepEqual(received, expected) {
		t.Error("Expected ", expected, " but received ", received)
	}
}

func TestMultiplyAllButItselfHaveZeroElement(t *testing.T) {
	values := multiplyallbutself.Integers{1, 2, 3, 4, 0, 6, 7, 0}
	expected := multiplyallbutself.Integers{0, 0, 0, 0, 0, 0, 0, 0}

	received := multiplyallbutself.MultiplyAllButSelf(values)

	if !reflect.DeepEqual(received, expected) {
		t.Error("Expected ", expected, " but received ", received)
	}
}

func TestMultiplyAllButItselfHaveSingleZeroElement(t *testing.T) {
	values := multiplyallbutself.Integers{1, 2, 3, 4, 5, 6, 7, 0}
	expected := multiplyallbutself.Integers{0, 0, 0, 0, 0, 0, 0, 5040}

	received := multiplyallbutself.MultiplyAllButSelf(values)

	if !reflect.DeepEqual(received, expected) {
		t.Error("Expected ", expected, " but received ", received)
	}
}

func TestMultiplyAllButItselfEmptyArray(t *testing.T) {
	values := multiplyallbutself.Integers{}
	expected := multiplyallbutself.Integers(nil)

	received := multiplyallbutself.MultiplyAllButSelf(values)

	if !reflect.DeepEqual(received, expected) {
		t.Error("Expected ", expected, " but received ", received)
	}
}

func TestMultiplyAllButItselfArrayWithSingleElement(t *testing.T) {
	values := multiplyallbutself.Integers{12}
	expected := multiplyallbutself.Integers{0}

	received := multiplyallbutself.MultiplyAllButSelf(values)

	if !reflect.DeepEqual(received, expected) {
		t.Error("Expected ", expected, " but received ", received)
	}
}
