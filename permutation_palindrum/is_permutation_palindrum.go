package permutationpalindrum

import "unicode"

func is_permutation_palindrum(word string) bool{
	char_map := map[rune]bool{}

	for _,char := range word{
		if !unicode.IsSpace(char){
			_, ok := char_map[char]
			if ok{
				delete(char_map,char)
			}else{
				char_map[char] = true
			}
		}
	}
	return len(char_map) <= 1
}