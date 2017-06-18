package substring

func preprocess(word string) []int {
	arr := []int{-1}
	j := 0
	for i := 1; i < len(word); i++ {
		if word[i] == word[j] {
			arr = append(arr, j)
			j++
		} else {
			j = j - 1
			if j != -1 {
				for j = arr[j]; j != -1 && word[i] != word[j]; {
					if j = j - 1; j != -1 {
						j = arr[j]
					}
				}
			}
			arr = append(arr, j)
			j++
		}
	}

	return arr
}

// KMP Knuth, Morris and Pratt string matching
// algorithm for finding the first index of a
// substring
func KMP(text, word string) int {
	if len(text) == 0 || len(word) == 0 {
		return -1
	}

	arr := preprocess(word)
	j := 0
	for i := 0; i < len(text)-(len(word)-j)+1; {
		k := 0
		for ; j < len(word) && text[i+k] == word[j]; j, k = j+1, k+1 {
		}
		if j == len(word) {
			return i + k - j
		}
		if j == 0 {
			i++
			j++
		} else {
			i = i + k
			j = arr[j-1] + 1
		}
	}
	return -1
}
