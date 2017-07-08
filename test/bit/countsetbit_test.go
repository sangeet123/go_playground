package bittest

import (
	"go_playground/bit"
	"testing"
)

func TestCountBitSet42(t *testing.T) {
	if bit.CountBitSet(42) != 3 {
		t.Error("Expected 3 but received ", bit.CountBitSet)
	}
}

func TestCountBitSet0(t *testing.T) {
	if bit.CountBitSet(0) != 0 {
		t.Error("Expected 0 but received ", bit.CountBitSet)
	}
}
