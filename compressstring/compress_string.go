package compressstring

import (
	"bytes"
	"fmt"
	"unicode/utf8"
)

// This is O(2n) solution
// This works for utf8 encoding scheme
func Compress(word string) string {
	char_map := map[rune]int{}

	for _, val := range word {
		char_map[val]++
	}

	var compressed bytes.Buffer
	for _, val := range word {
		if count, ok := char_map[val]; ok {
			compressed.WriteString(string(val))
			fmt.Fprintf(&compressed, "%d", count)
		}
		delete(char_map, val)
	}

	compressedString := compressed.String()
	if len(compressedString) < len(word) {
		return compressedString
	}

	return word
}

// can be done in a single pass using following approach
func CompressEfficient(word string) string {
	if len(word) < 2 {
		return word
	}

	count := 1
	var compressed bytes.Buffer
	prevChar, width := utf8.DecodeRuneInString(word[0:])
	for _, val := range word[width:] {
		if val != prevChar {
			compressed.WriteString(string(prevChar))
			fmt.Fprintf(&compressed, "%d", count)
			prevChar = val
			count = 1
		} else {
			count++
		}
	}

	compressed.WriteString(string(prevChar))
	fmt.Fprintf(&compressed, "%d", count)
	compressedString := compressed.String()

	if len(compressedString) < len(word) {
		return compressedString
	}

	return word
}
