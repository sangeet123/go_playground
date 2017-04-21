package uniquechar

import "testing"

func TestStringHavingUniqueChar(t *testing.T) {
	if !IsUnique("abcdef") {
		t.Error("Expected true, got ", false)
	}
}

func TestStringNotHavingUniqueChar(t *testing.T) {
	if IsUnique("Hello") {
		t.Error("Expected false, got", true)
	}
}

func TestEmptyString(t *testing.T) {
	if !IsUnique("") {
		t.Error("Expected true, got", false)
	}
}

func TestOneCharString(t *testing.T) {
	if !IsUnique("a") {
		t.Error("Expected true, got", false)
	}
}
