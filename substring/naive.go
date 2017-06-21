package substring

// Naive returns start first index of word and true
// a word is substring of text else return -1 and false
func Naive(text, word string) int {
	if failEarly(text, word) {
		return -1
	}
	naive := func(text, word string) int {
		for i := 0; i < len(text)-len(word)+1; i++ {
			if word == text[i:i+len(word)] {
				return i
			}
		}
		return -1
	}
	return naive(text, word)
}
