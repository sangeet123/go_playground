package substring

func failEarly(text, word string) bool {
	return len(text) == 0 || len(word) == 0 || len(text) < len(word)
}
