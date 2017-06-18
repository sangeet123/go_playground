package substring

// Naive returns start first index of word and true
// a word is substring of text else return -1 and false
func Naive(text, word string) int {

	if len(text) == 0 || len(word) == 0 {
		return -1
	}

	for i := 0; i < len(text)-len(word)+1; i++ {
		j := 0
		for ; j < len(word) && word[j] == text[i+j]; j++ {
		}
		if j == len(word) {
			return i
		}
	}
	return -1
}
