package arrayprobtest

import (
	"go_playground/arrayprob"
	"testing"
)

func TestValidRotatedWordExample1(t *testing.T) {
	word := "mangobanana"
	rotated := "bananamango"
	if !arrayprob.IsRotated(word, rotated) {
		t.Error("Expected true but returned false")
	}
}

func TestValidRotatedWordExample2(t *testing.T) {
	word := "mangobanana"
	rotated := "nanamangoba"
	if !arrayprob.IsRotated(word, rotated) {
		t.Error("Expected true but returned false")
	}
}

func TestInValidRotatedWordUnEqualLength(t *testing.T) {
	word := "mangobana"
	rotated := "bana"
	if arrayprob.IsRotated(word, rotated) {
		t.Error("Expected false but returned true")
	}
}

func TestValidBothEmptyStrings(t *testing.T) {
	word := ""
	rotated := ""
	if !arrayprob.IsRotated(word, rotated) {
		t.Error("Expected false but returned true")
	}
}
