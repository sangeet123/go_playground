package permutationpalindrum

import "unicode"

// Algorithm:
// 1: Create a map
// loop through all the char in words
// if char not found insert into map
// if char found delete that char
// After the loop ends check if map has zero or singl entry
func is_permutation_palindrum(word string) bool {
	char_map := map[rune]bool{}

	for _, char := range word {
		if !unicode.IsSpace(char) {
			_, ok := char_map[char]
			if ok {
				delete(char_map, char)
			} else {
				char_map[char] = true
			}
		}
	}
	return len(char_map) <= 1
}
