package permutationchecktest

import (
	"go_playground/permutationcheck"
	"testing"
)


func TestValidPermutationStrings(t *testing.T) {
	if !permutationcheck.CheckPermutation("Hello", "oleHl") {
		t.Error("Expected true, but got", false)
	}
}

func TestValidPermutationStringsWithUpperAndLowerCaseCharacters(t *testing.T) {
	if !permutationcheck.CheckPermutation("HELLO", "olehl") {
		t.Error("Expected true, but got", false)
	}
}

func TestInvalidValidPermutationStrings(t *testing.T) {
	if permutationcheck.CheckPermutation("HELLLO", "oloehl") {
		t.Error("Expected true, but got", false)
	}
}

func TestValidPermutationEmptyStrings(t *testing.T) {
	if !permutationcheck.CheckPermutation("", "") {
		t.Error("Expected true, but got", false)
	}
}

func TestInValidPermutationUnEqualLengthStrings(t *testing.T) {
	if permutationcheck.CheckPermutation("HELLO", "olehll") {
		t.Error("Expected true, but got", false)
	}
}
