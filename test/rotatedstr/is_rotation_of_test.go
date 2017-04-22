package rotatedstrtest

import (
	"go_playground/rotatedstr"
	"testing"
)

func TestValidRotatedWordExample1(t *testing.T) {
	word := "mangobanana"
	rotated := "bananamango"
	if !rotatedstr.IsRotated(word, rotated) {
		t.Error("Expected true but returned false")
	}
}

func TestValidRotatedWordExample2(t *testing.T) {
	word := "mangobanana"
	rotated := "nanamangoba"
	if !rotatedstr.IsRotated(word, rotated) {
		t.Error("Expected true but returned false")
	}
}

func TestInValidRotatedWordUnEqualLength(t *testing.T) {
	word := "mangobana"
	rotated := "bana"
	if rotatedstr.IsRotated(word, rotated) {
		t.Error("Expected false but returned true")
	}
}

func TestValidBothEmptyStrings(t *testing.T) {
	word := ""
	rotated := ""
	if !rotatedstr.IsRotated(word, rotated) {
		t.Error("Expected false but returned true")
	}
}
