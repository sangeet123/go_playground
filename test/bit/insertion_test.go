package bittest

import (
	"go_playground/bit"
	"testing"
)

func TestInsertion(t *testing.T) {
	if val := bit.Insert(1024, 19, 6, 2); val != 1100 {
		t.Error("Expected 1100 but received ", val)
	}
}

func TestInsertionRightIsZero(t *testing.T) {
	if val := bit.Insert(1156, 19, 4, 0); val != 1171 {
		t.Error("Expected 1100 but received ", val)
	}
}
