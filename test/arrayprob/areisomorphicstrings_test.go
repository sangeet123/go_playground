package arrayprobtest

import (
	"go_playground/arrayprob"
	"testing"
)

func TestTwoIsoMorphicString(t *testing.T) {
	word1 := "egg"
	word2 := "add"
	if !arrayprob.AreIsomorphicStrings(word1, word2) {
		t.Error("Expected true but returned false")
	}
}

func TestTwoIsoNonMorphicStringExample1(t *testing.T) {
	word1 := "adge"
	word2 := "agfk"
	if arrayprob.AreIsomorphicStrings(word1, word2) {
		t.Error("Expected false but returned true")
	}
}

func TestTwoIsoNonMorphicStringDoubleMappingOfSingleChar(t *testing.T) {
	word1 := "foo"
	word2 := "bar"
	if arrayprob.AreIsomorphicStrings(word1, word2) {
		t.Error("Expected false but returned true")
	}
}

func TestTwoIsoNonMorphicStringExample2(t *testing.T) {
	word1 := "ab"
	word2 := "aa"
	if arrayprob.AreIsomorphicStrings(word1, word2) {
		t.Error("Expected false but returned true")
	}
}
