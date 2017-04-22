package oneeditaway

import (
	"go_playground/oneeditaway"
	"testing"
)

func TestValidOneWayEditDelete(t *testing.T) {
	if !oneeditaway.IsOneEditAway("pale", "ple") {
		t.Error("Expected true but got false")
	}
}

func TestValidOneWayEditSubstitute(t *testing.T) {
	if !oneeditaway.IsOneEditAway("ale", "tle") {
		t.Error("Expected true but got false")
	}
}

func TestValidOneWayEditEqualString(t *testing.T) {
	if !oneeditaway.IsOneEditAway("word", "word") {
		t.Error("Expected true but got false")
	}
}

func TestValidOneWayEditInsertion(t *testing.T) {
	if !oneeditaway.IsOneEditAway("ple", "pale") {
		t.Error("Expected true but got false")
	}
}

func TestValidEmptyStrings(t *testing.T) {
	if !oneeditaway.IsOneEditAway("", "") {
		t.Error("Expected true but got false")
	}
}

func TestValidOneCharAndEmptyStrings(t *testing.T) {
	if !oneeditaway.IsOneEditAway("", "a") {
		t.Error("Expected true but got false")
	}
}

func TestValidEmptyStringsAndOneChar(t *testing.T) {
	if !oneeditaway.IsOneEditAway("a", "") {
		t.Error("Expected true but got false")
	}
}

func TestInvalidTwoCharDiffString(t *testing.T) {
	if oneeditaway.IsOneEditAway("abc", "adb") {
		t.Error("Expected false but got true")
	}
}

func TestInvalidUnEqualLenghtString(t *testing.T) {
	if oneeditaway.IsOneEditAway("a", "adb") {
		t.Error("Expected false but got true")
	}
}

//test cases from cracking the coding interview
func TestValidStrings_cracking_example_1(t *testing.T) {
	if !oneeditaway.IsOneEditAway("pales", "pale") {
		t.Error("Expected true but got false")
	}
}

func TestValidStrings_cracking_example_2(t *testing.T) {
	if !oneeditaway.IsOneEditAway("pale", "bale") {
		t.Error("Expected true but got false")
	}
}

func TestInValidStrings_cracking_example_3(t *testing.T) {
	if oneeditaway.IsOneEditAway("pale", "bake") {
		t.Error("Expected true but got false")
	}
}
