package oneeditaway

func check_if_convertible(word1 string, word2 string) bool {
	i := 0
	for i < len(word2) && word1[i] == word2[i] {
		i++
	}
	if i == len(word2) {
		return true
	}
	return word1[1+1:] == word2[i:]
}

//char sensitive
func IsOneEditAway(src string, dest string) bool {

	length_diff := len(src) - len(dest)

	if !(length_diff == -1 || length_diff == 0 || length_diff == 1) {
		return false
	}

	if length_diff == 1 {
		return check_if_convertible(src, dest)
	}

	if length_diff == -1 {
		return check_if_convertible(dest, src)
	}

	no_of_edits := 0
	for i := 0; i < len(src) && no_of_edits <= 1; i++ {
		if src[i] != dest[i] {
			no_of_edits++
		}
	}

	return no_of_edits <= 1
}
