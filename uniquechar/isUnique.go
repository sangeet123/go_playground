package uniquechar

func IsUnique(word string) bool {
	charMap := make(map[rune]struct{})
	for _, val := range word {
		if _, ok := charMap[val]; ok {
			return false
		}
		charMap[val] = struct{}{}
	}
	return true
}
