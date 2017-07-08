package bittest

import (
	"go_playground/bit"
	"testing"
)

func TestConversionForAllBitUnequal(t *testing.T) {
	if bit.Conversion(42, 21) != 6 {
		t.Error("Expected 6 but received ", bit.Conversion(42, 21))
	}
}

func TestConversionForAllBitEqual(t *testing.T) {
	if bit.Conversion(10, 10) != 0 {
		t.Error("Expected 0 but received ", bit.Conversion(10, 10))
	}
}

func TestConversionExample1(t *testing.T) {
	if bit.Conversion(7, 1) != 2 {
		t.Error("Expected 2 but received ", bit.Conversion(7, 1))
	}
}
