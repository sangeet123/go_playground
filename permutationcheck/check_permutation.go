package permutationcheck

import "unicode"

func CheckPermutation(src, dest string) bool {

	if len(src) != len(dest) {
		return false
	}

	charMap := make(map[rune]int)

	for _, char := range src {
		lowerCaseChar := unicode.ToLower(char)
		charMap[lowerCaseChar]++
	}

	for _, char := range dest {
		lowerCaseChar := unicode.ToLower(char)
		if _, ok := charMap[lowerCaseChar]; !ok || charMap[lowerCaseChar] == 0 {
			return false
		}
		charMap[lowerCaseChar]--
	}
	return true
}
