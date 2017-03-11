package compressstring

import "testing"

func TestCompressableStringEfficient(t *testing.T){
	test_word := "aaabbb"
	expected := "a3b3"
	received := compress(test_word)

	if compress_efficient(test_word) != expected{
		t.Error("Expected ", expected, " but got", received)
	}
}

func TestNonCompressStringEfficient(t *testing.T){
	test_word := "ab"
	expected := "ab"
	received := compress(test_word)

	if compress_efficient(test_word) != expected{
		t.Error("Expected ", expected, "but got ", received)
	}
}

func TestCompressableStringLessEfficient(t *testing.T){
	test_word := "aaabbb"
	expected := "a3b3"
	received := compress(test_word)

	if compress(test_word) != expected{
		t.Error("Expected ", expected, " but got", received)
	}
}
