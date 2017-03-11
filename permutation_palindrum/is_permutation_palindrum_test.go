package permutationpalindrum

import "testing"

func TestValidPalindrumWordWithoutSpace(t *testing.T){
	word := "ababadd"
	if !is_permutation_palindrum(word){
		t.Error("should have been true but got false")
	}
}

func TestValidPalindrumWordWithSpace(t *testing.T){
	word := "\n\taba ba\n\t dd\t\n"
	if !is_permutation_palindrum(word){
		t.Error("should have been true but got false")
	}
}

func TestInValidPalindrumWord(t *testing.T){
	word := `adadef`
	if is_permutation_palindrum(word){
		t.Error("should have been false but got true")
	}
}