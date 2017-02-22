package permutationcheck

import "unicode"

func CheckPermutation(src, dest string) bool{

	if len(src) != len(dest){
		return false
	}

	charMap := make(map[rune]int)

	for _, char := range src {
		lowerCaseChar := unicode.ToLower(char)
		val, ok := charMap[lowerCaseChar]
		if ok {
			charMap[lowerCaseChar] = val + 1
		}else{
			charMap[lowerCaseChar] = 1
		}
	}

	for _, char := range dest {
		lowerCaseChar := unicode.ToLower(char)
		val, ok := charMap[lowerCaseChar]
		if !ok || val == 0 {
			return false
		}else{
			charMap[lowerCaseChar] = val - 1
		}
	}
	return true
}