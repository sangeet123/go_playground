package multiplyallbutitself

import "testing"
import "reflect"

func TestMultiplyAllButItself(t *testing.T) {
	values := Integers{1, 2, 3, 4, 5, 6, 7}
	expected := Integers{5040, 2520, 1680, 1260, 1008, 840, 720}

	received := multiply_all_but_itself(values)

	if !reflect.DeepEqual(received, expected) {
		t.Error("Expected ", expected, " but received ", received)
	}
}

func TestMultiplyAllButItselfHaveZeroElement(t *testing.T) {
	values := Integers{1, 2, 3, 4, 0, 6, 7, 0}
	expected := Integers{0, 0, 0, 0, 0, 0, 0, 0}

	received := multiply_all_but_itself(values)

	if !reflect.DeepEqual(received, expected) {
		t.Error("Expected ", expected, " but received ", received)
	}
}

func TestMultiplyAllButItselfHaveSingleZeroElement(t *testing.T) {
	values := Integers{1, 2, 3, 4, 5, 6, 7, 0}
	expected := Integers{0, 0, 0, 0, 0, 0, 0, 5040}

	received := multiply_all_but_itself(values)

	if !reflect.DeepEqual(received, expected) {
		t.Error("Expected ", expected, " but received ", received)
	}
}

func TestMultiplyAllButItselfEmptyArray(t *testing.T) {
	values := Integers{}
	expected := Integers(nil)

	received := multiply_all_but_itself(values)

	if !reflect.DeepEqual(received, expected) {
		t.Error("Expected ", expected, " but received ", received)
	}
}

func TestMultiplyAllButItselfArrayWithSingleElement(t *testing.T) {
	values := Integers{12}
	expected := Integers{0}

	received := multiply_all_but_itself(values)

	if !reflect.DeepEqual(received, expected) {
		t.Error("Expected ", expected, " but received ", received)
	}
}
