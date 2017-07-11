package bittest

import (
	"go_playground/bit"
	"testing"
)

func TestSwapEvenOddBitExample1(t *testing.T) {
	if val := bit.SwapEvenOddBit(int32(1024)); val != int32(2048) {
		t.Error("Expected 1100 but received ", val)
	}
}

func TestSwapEvenOddBitExample2(t *testing.T) {
	if val := bit.SwapEvenOddBit(int32(42)); val != int32(21) {
		t.Error("Expected 1100 but received ", val)
	}
}

func TestSwapEvenOddBitExample3(t *testing.T) {
	if val := bit.SwapEvenOddBit(int32(-4195461)); val != int32(-8390729) {
		t.Error("Expected 1100 but received ", val)
	}
}
