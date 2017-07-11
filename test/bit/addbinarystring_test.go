package bittest

import (
	"go_playground/bit"
	"testing"
)

func TestConversionForAllBitUnequal(t *testing.T) {
	o1 := []string{"111111", "111111", "1011111", "101", "1", "1", "100000"}
	o2 := []string{"101", "1", "1", "111111", "111111", "1011111", ""}
	rs := []string{"1000100", "1000000", "1100000", "1000100", "1000000", "1100000", "100000"}

	for i := 0; i < len(o1); i++ {
		if r := bit.Add(o1[i], o2[i]); r != rs[i] {
			t.Error("expected ", rs[i], " but received ", r)
		}
	}
}
