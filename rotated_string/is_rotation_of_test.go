package rotated

import "testing"

func TestValidRotatedWordExample1(t *testing.T) {
	word := "mangobanana"
	rotated := "bananamango"
	if !is_rotated_of(word, rotated) {
		t.Error("Expected true but returned false")
	}
}

func TestValidRotatedWordExample2(t *testing.T) {
	word := "mangobanana"
	rotated := "nanamangoba"
	if !is_rotated_of(word, rotated) {
		t.Error("Expected true but returned false")
	}
}

func TestInValidRotatedWordUnEqualLength(t *testing.T) {
	word := "mangobana"
	rotated := "bana"
	if is_rotated_of(word, rotated) {
		t.Error("Expected false but returned true")
	}
}

func TestValidBothEmptyStrings(t *testing.T) {
	word := ""
	rotated := ""
	if !is_rotated_of(word, rotated) {
		t.Error("Expected false but returned true")
	}
}
