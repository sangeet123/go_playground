package permutationcheck

import "testing"

func TestValidPermutationStrings(t *testing.T){
	if !CheckPermutation("Hello", "oleHl"){
		t.Error("Expected true, but got", false)
	}
}

func TestValidPermutationStringsWithUpperAndLowerCaseCharacters(t *testing.T){
	if !CheckPermutation("HELLO", "olehl"){
		t.Error("Expected true, but got", false)
	}
}

func TestInvalidValidPermutationStrings(t *testing.T){
	if CheckPermutation("HELLLO", "oloehl"){
		t.Error("Expected true, but got", false)
	}
}

func TestValidPermutationEmptyStrings(t *testing.T){
	if !CheckPermutation("", ""){
		t.Error("Expected true, but got", false)
	}
}

func TestInValidPermutationUnEqualLengthStrings(t *testing.T){
	if CheckPermutation("HELLO", "olehll"){
		t.Error("Expected true, but got", false)
	}
}