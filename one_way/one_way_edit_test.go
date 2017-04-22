package oneeditaway

import "testing"

func TestValidOneWayEditDelete(t *testing.T) {
	if !is_one_edit_away("pale", "ple") {
		t.Error("Expected true but got false")
	}
}

func TestValidOneWayEditSubstitute(t *testing.T) {
	if !is_one_edit_away("ale", "tle") {
		t.Error("Expected true but got false")
	}
}

func TestValidOneWayEditEqualString(t *testing.T) {
	if !is_one_edit_away("word", "word") {
		t.Error("Expected true but got false")
	}
}

func TestValidOneWayEditInsertion(t *testing.T) {
	if !is_one_edit_away("ple", "pale") {
		t.Error("Expected true but got false")
	}
}

func TestValidEmptyStrings(t *testing.T) {
	if !is_one_edit_away("", "") {
		t.Error("Expected true but got false")
	}
}

func TestValidOneCharAndEmptyStrings(t *testing.T) {
	if !is_one_edit_away("", "a") {
		t.Error("Expected true but got false")
	}
}

func TestValidEmptyStringsAndOneChar(t *testing.T) {
	if !is_one_edit_away("a", "") {
		t.Error("Expected true but got false")
	}
}

func TestInvalidTwoCharDiffString(t *testing.T) {
	if is_one_edit_away("abc", "adb") {
		t.Error("Expected false but got true")
	}
}

func TestInvalidUnEqualLenghtString(t *testing.T) {
	if is_one_edit_away("a", "adb") {
		t.Error("Expected false but got true")
	}
}

//test cases from cracking the coding interview
func TestValidStrings_cracking_example_1(t *testing.T) {
	if !is_one_edit_away("pales", "pale") {
		t.Error("Expected true but got false")
	}
}

func TestValidStrings_cracking_example_2(t *testing.T) {
	if !is_one_edit_away("pale", "bale") {
		t.Error("Expected true but got false")
	}
}

func TestInValidStrings_cracking_example_3(t *testing.T) {
	if is_one_edit_away("pale", "bake") {
		t.Error("Expected true but got false")
	}
}
