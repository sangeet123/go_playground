package uniquechar

func IsUnique(word string) bool {
	charMap := make(map[rune]bool)
	for _, val := range word {
		if _, ok := charMap[val]; ok {
			return false
		}
		charMap[val] = true
	}
	return true
}
