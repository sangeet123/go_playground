package compressstring

import "strconv"
import "fmt"

// This is O(2n) solution
func compress(word string)string{
	char_map := map[rune]int{}

	for _, val := range word{
		_, ok := char_map[val]

		if !ok{
			char_map[val] = 1
		}else{
			char_map[val] = char_map[val] + 1
		}
	}

	compressed := ""
	for _, val := range word{
		_, ok := char_map[val]
		if ok{
			compressed = compressed + string(val) + strconv.Itoa(char_map[val])
		}
		delete(char_map,val)
	}

	if len(compressed) < len(word){
		return compressed
	}

	return word
}

// can be done in a single pass using following approach
func compress_efficient(word string)string{
	if len(word) < 2{
		return word
	}

	count := 1
	compressed := ""
	for i := 1 ; i < len(word) && len(compressed) < len(word) ; i++{
		if word[i] != word[i-1]{
			compressed = string(word[i-1]) + strconv.Itoa(count)
			count = 1
		}else{
			count++
		}
	}

	compressed = compressed + string(word[len(word)-1]) + strconv.Itoa(count)
	fmt.Println(compressed)

	if len(compressed) < len(word){
		return compressed
	}

	return word
}
