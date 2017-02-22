package uniquechar

func IsUnique(word string) bool{
	charMap := make(map[rune]bool)
	for _, val := range word {
		 _, ok := charMap[val]
		 if(ok){
		 	return false
		 }
		 charMap[val] = true
	}
	return true
}
