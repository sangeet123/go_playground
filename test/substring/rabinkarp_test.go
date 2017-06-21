package naivetest

import (
	"go_playground/substring"
	"testing"
)

func TestIsASubString(t *testing.T) {
	text := "mangobanana"
	word := "banana"
	if i := substring.Naive(text, word); i != 5 {
		t.Error("expected ", 5, " but received ", i)
	}
}

func TestIsNotASubString(t *testing.T) {
	text := "mangobanana"
	word := "bananaaaa"
	if i := substring.Naive(text, word); i != -1 {
		t.Error("expected ", -1, " but received ", i)
	}
}

func TestTextAndWordOfEqualLength(t *testing.T) {
	text := "banana"
	word := "banana"
	if i := substring.Naive(text, word); i != 0 {
		t.Error("expected ", -1, " but received ", i)
	}
}

func TestTextAndWordBothEmpty(t *testing.T) {
	text := ""
	word := ""
	if i := substring.Naive(text, word); i != -1 {
		t.Error("expected ", -1, " but received ", i)
	}
}

func TestTextEmpty(t *testing.T) {
	text := ""
	word := "banana"
	if i := substring.Naive(text, word); i != -1 {
		t.Error("expected ", -1, " but received ", i)
	}
}

func TestWordEmpty(t *testing.T) {
	text := "banana"
	word := ""
	if i := substring.Naive(text, word); i != -1 {
		t.Error("expected ", -1, " but received ", i)
	}
}

func TestIsASubStringLongText(t *testing.T) {
	text := "aaabbaaabbaaabbaaaaaaaabbbaaaabaaabb"
	word := "aaaabbbaaaabaaabb"
	if i := substring.Naive(text, word); i != 19 {
		t.Error("expected ", 18, " but received ", i)
	}
}
