package compressstringtest

import (
	"go_playground/compressstring"
	"testing"
)


func TestCompressableStringEfficient(t *testing.T) {
	test_word := "aaabbb"
	expected := "a3b3"
	received := compressstring.CompressEfficient(test_word)

	if received != expected {
		t.Error("Expected ", expected, " but got", received)
	}
}

func TestNonCompressStringEfficient(t *testing.T) {
	test_word := "ab"
	expected := "ab"
	received := compressstring.CompressEfficient(test_word)

	if received != expected {
		t.Error("Expected ", expected, "but got ", received)
	}
}

func TestEmptyStringEfficient(t *testing.T) {
	test_word := ""
	expected := ""
	received := compressstring.CompressEfficient(test_word)

	if received != expected {
		t.Error("Expected ", expected, "but got ", received)
	}
}

func TestCompressableStringLessEfficient(t *testing.T) {
	test_word := "aaabbb"
	expected := "a3b3"
	received := compressstring.Compress(test_word)

	if received != expected {
		t.Error("Expected ", expected, " but got", received)
	}
}

func TestNonCompressStringLessEfficient(t *testing.T) {
	test_word := "ab"
	expected := "ab"
	received := compressstring.Compress(test_word)

	if received != expected {
		t.Error("Expected ", expected, "but got ", received)
	}
}
