package uniquechartest

import (
	"go_playground/uniquechar"
	"testing"
)

func TestStringHavingUniqueChar(t *testing.T) {
	if !uniquechar.IsUnique("abcdef") {
		t.Error("Expected true, got ", false)
	}
}

func TestStringNotHavingUniqueChar(t *testing.T) {
	if uniquechar.IsUnique("Hello") {
		t.Error("Expected false, got", true)
	}
}

func TestEmptyString(t *testing.T) {
	if !uniquechar.IsUnique("") {
		t.Error("Expected true, got", false)
	}
}

func TestOneCharString(t *testing.T) {
	if !uniquechar.IsUnique("a") {
		t.Error("Expected true, got", false)
	}
}
