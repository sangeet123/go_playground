package kmptest

import (
	"go_playground/substring"
	"testing"
)

func TestIsASubString(t *testing.T) {
	text := "aaaab"
	word := "aaab"
	if i := substring.KMP(text, word); i != 1 {
		t.Error("expected ", 5, " but received ", i)
	}
}

func TestASubStringLongTextAndLongWord(t *testing.T) {
	text := "aaabbaaabbaaabbaaaaaaaabbbaaaabaaabb"
	word := "aaaabbbaaaabaaabb"
	if i := substring.KMP(text, word); i != 19 {
		t.Error("expected ", 18, " but received ", i)
	}
}

func TestIsNotASubString(t *testing.T) {
	text := "mangobanana"
	word := "bananaaaa"
	if i := substring.KMP(text, word); i != -1 {
		t.Error("expected ", -1, " but received ", i)
	}
}

func TestTextAndWordOfEqualLength(t *testing.T) {
	text := "banana"
	word := "banana"
	if i := substring.KMP(text, word); i != 0 {
		t.Error("expected ", -1, " but received ", i)
	}
}
