package permutationpalindrum

import (
	"go_playground/permutationpalindrum"
	"testing"
)

func TestValidPalindrumWordWithoutSpace(t *testing.T) {
	word := "ababadd"
	if !permutationpalindrum.IsPermutationPalindrum(word) {
		t.Error("should have been true but got false")
	}
}

func TestValidPalindrumWordWithSpace(t *testing.T) {
	word := "\n\taba ba\n\t dd\t\n"
	if !permutationpalindrum.IsPermutationPalindrum(word) {
		t.Error("should have been true but got false")
	}
}

func TestInValidPalindrumWord(t *testing.T) {
	word := `adadef`
	if permutationpalindrum.IsPermutationPalindrum(word) {
		t.Error("should have been false but got true")
	}
}
