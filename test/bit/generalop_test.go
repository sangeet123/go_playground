package bittest

import (
	"go_playground/bit"
	"testing"
)

func TestIsSet(t *testing.T) {
	expected := []bool{false, true, false, true, false, true, false}
	for i, val := range expected {
		p := uint(i)
		if val != bit.IsSet(42, p) {
			t.Error("Expected ", expected[i], " but received ", bit.IsSet(42, p))
		}
	}
}

func TestSetBit(t *testing.T) {
	positionsToSet := []uint{1, 3, 5}
	for _, val := range positionsToSet {
		setNo := bit.SetBit(42, val)
		if !bit.IsSet(setNo, val) {
			t.Error("Expected true but returned false")
		}
	}
}

func TestGetBit(t *testing.T) {
	expectedBits := []int{0, 1, 0, 1, 0, 1}
	for i, val := range expectedBits {
		p := uint(i)
		if bit.GetBit(42, p) != val {
			t.Error("Expected ", val, " but returned ", bit.GetBit(42, p))
		}
	}
}

func TestClearBit(t *testing.T) {
	positionsToClear := []uint{1, 3, 5}
	for _, val := range positionsToClear {
		clearNo := bit.ClearBit(42, val)
		if bit.IsSet(clearNo, val) {
			t.Error("Expected false but returned true")
		}
	}
}

func TestUpdateBit(t *testing.T) {
	updateTo := []bool{true, false, true, false, true, false}
	no := 42
	for i, val := range updateTo {
		p := uint(i)
		no = bit.UpdateBit(no, p, val)
	}
	if no != 21 {
		t.Error("Expected 21 but received ", no)
	}
}
