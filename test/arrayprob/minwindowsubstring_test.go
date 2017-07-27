package arrayprobtest

import (
	"go_playground/arrayprob"
	"testing"
)

func TestGetMinWindowSubStringExample1(t *testing.T) {
	word1 := "xyzaxxxxyzabxyyyyzbcxyzzzzaaaaa"
	word2 := "abc"
	if val := arrayprob.GetMinWindowSubString(word1, word2); "bcxyzzzza" != val {
		t.Error("Expected bcxyzzzza but returned ", val)
	}
}

func TestGetMinWindowSubStringExample2(t *testing.T) {
	word1 := "xyzaxxxxyzabxyyyyzbcxyzzzzaaaaaaaaab"
	word2 := "abbc"
	if val := arrayprob.GetMinWindowSubString(word1, word2); "abxyyyyzbc" != val {
		t.Error("Expected abxyyyyzbc but returned ", val)
	}
}

func TestGetMinWindowSubStringExample3(t *testing.T) {
	word1 := "xyzaxxxxyzabxyyyyzbcxyzzzzaaaaaaaababc"
	word2 := "abbc"
	if val := arrayprob.GetMinWindowSubString(word1, word2); "babc" != val {
		t.Error("Expected babc but returned ", val)
	}
}

func TestGetMinWindowSubStringExample4(t *testing.T) {
	word1 := "xyzaxxxxyzabxyyyyzbcxyzzzzaaaaaaaababc"
	word2 := "ababcc"
	if val := arrayprob.GetMinWindowSubString(word1, word2); "cxyzzzzaaaaaaaababc" != val {
		t.Error("Expected babc but returned ", val)
	}
}
