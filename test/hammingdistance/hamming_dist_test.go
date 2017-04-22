package hammingdisttest

import (
	"go_playground/hammingdist"
	"testing"
)

func TestValidDistSmallNumber(t *testing.T) {
	dist := hammingdist.GetHammingDistance(1, 4)
	if dist != 2 {
		t.Error("Expected 2, but got", dist)
	}
}

func TestValidDistMidValNumber(t *testing.T) {
	dist := hammingdist.GetHammingDistance(1234, 4567)
	// 1234 => 0010011010010
	// 4567 => 1000111010111
	if dist != 5 {
		t.Error("Expected 2, but got", dist)
	}
}

func TestValidDistLargeNumber(t *testing.T) {
	dist := hammingdist.GetHammingDistance(4444445555, 66666655555)
	// 4444445555  => 000100001000111010001101101101110011
	// 66666655555 => 111110000101101001000110111101000011
	if dist != 16 {
		t.Error("Expected 2, but got", dist)
	}
}

func TestValidDistEqualNumber(t *testing.T) {
	dist := hammingdist.GetHammingDistance(1234, 1234)
	if dist != 0 {
		t.Error("Expected 2, but got", dist)
	}
}
