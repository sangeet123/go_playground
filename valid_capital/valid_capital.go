package validcapital

import "unicode"

// Check if word is valid word or not
// if a word starts with capital letter followed by all small case letters it is a valid word eg "Hello"
// if a word has all upper case letter it is a valid word eg "USA"
// if a word has all small case letter it is a valid word eg "hello"
// words with leading or trailing whitespaces or has whitespaces in between characters are not valid eg "h ello"
// the credit for this question goes to leetcode
func ValidCapital(word string) bool{
	
	if len(word) == 0{
		return true
	}

	allUpper := func(word string) (bool, int) {
		for i, char := range word{
			if !unicode.IsUpper(char){
				return false, i
			}
		}
		return true,-1
	}

	allLower := func(word string) (bool, int) {
		for i, char := range word{
			if !unicode.IsLower(char){
				return false, i
			}
		}

		return true, -1
	}

	firstCharUpper := func(word string) (bool, int) {
		if(!unicode.IsUpper(rune(word[0]))){
			return false, 0
		}
		ok, index := allLower(word[1:])
		return ok, index + 1
	}

	isValid, index := firstCharUpper(word)

	if(isValid){
		return true
	}

	if(index == 1){
		isValid, _ = allUpper(word)
		return true		
	}

	if (index == 0){
		isValid, _ = allLower(word)
		return true
	}

	return false
}