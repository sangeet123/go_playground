package bittest

import (
	"go_playground/bit"
	"testing"
)

//10101010
func TestFlipBitToWinExample1(t *testing.T) {
	if l := bit.GetLongFlipLen(170); l != 3 {
		t.Error("Expected 6 but received ", l)
	}
}

//11110101011
func TestFlipBitToWinExample2(t *testing.T) {
	if l := bit.GetLongFlipLen(1963); l != 6 {
		t.Error("Expected 6 but received ", l)
	}
}

//111101110001110111
func TestFlipBitToWinExample3(t *testing.T) {
	if l := bit.GetLongFlipLen(253047); l != 8 {
		t.Error("Expected 6 but received ", l)
	}
}

//0
func TestFlipBitToWinExample4(t *testing.T) {
	if l := bit.GetLongFlipLen(0); l != 1 {
		t.Error("Expected 6 but received ", l)
	}
}

//11111
func TestFlipBitToWinExample5(t *testing.T) {
	if l := bit.GetLongFlipLen(31); l != 6 {
		t.Error("Expected 6 but received ", l)
	}
}

//1
func TestFlipBitToWinExample6(t *testing.T) {
	if l := bit.GetLongFlipLen(1); l != 2 {
		t.Error("Expected 6 but received ", l)
	}
}

//1111111111111111....64
func TestFlipBitToWinExample7(t *testing.T) {
	if l := bit.GetLongFlipLen(-1); l != 0 {
		t.Error("Expected 6 but received ", l)
	}
}

//1111111111111111....64
func TestFlipBitToWinExample8(t *testing.T) {
	if l := bit.GetLongFlipLen(-2); l != 63 {
		t.Error("Expected 6 but received ", l)
	}
}

//1111111111111111....64
func TestFlipBitToWinExample9(t *testing.T) {
	if l := bit.GetLongFlipLen(-3); l != 63 {
		t.Error("Expected 6 but received ", l)
	}
}
