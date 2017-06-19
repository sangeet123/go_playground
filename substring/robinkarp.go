package substring

const (
	d = 256
	p = 101
)

// RobinKarp for string matching
func RobinKarp(text, word string) int {
	if failEarly(text, word) {
		return -1
	}

	h := 1
	for i := 0; i < len(word); i++ {
		h = (h * d) % p
	}

	textHash := 0
	wordHash := 0

	for i := 0; i < len(word); i++ {
		textHash = (d*textHash + int(text[i])) % p
		wordHash = (d*wordHash + int(word[i])) % p
	}

	for i := 0; i < len(text)-len(word)+1; i++ {
		if textHash == wordHash && word == text[i:i+len(word)-1] {
			return i
		}
		if i+len(word) < len(text) {
			textHash = (d*textHash - h*int(text[i]) + int(text[i+len(word)])) % p
		}
	}
	return -1
}
