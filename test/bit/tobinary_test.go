package bittest

import (
	"go_playground/bit"
	"testing"
)

func TestToBinary(t *testing.T) {
	if val := bit.ToBinary(0.625); val != "0.101" {
		t.Error("Expected 1100 but received ", val)
	}

	if val := bit.ToBinary(0.75); val != "0.11" {
		t.Error("Expected 1100 but received ", val)
	}

	if val := bit.ToBinary(0.40625); val != "0.01101" {
		t.Error("Expected 1100 but received ", val)
	}
}
