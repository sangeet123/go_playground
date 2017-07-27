package arrayprob

func GetMinWindowSubString(str string, dest string) string {

	if len(str) < len(dest) {
		return ""
	}

	destMap := createMap(dest)
	strMap := make(map[rune]int)
	diff := createMap(dest)

	l := 0
	result := str
	for i, ch := range str {

		if _, ok := destMap[ch]; ok {
			strMap[ch]++
		}

		if _, ok := diff[ch]; ok {
			diff[ch]--
			if diff[ch] == 0 {
				delete(diff, ch)
			}
		}

		// if length of diff is equal to zero we found the window that has
		// all characters of dest in it
		if len(diff) == 0 {
			l += sqeezeFromLeft(str[l:i+1], destMap, strMap)

			// updating result if its length is already greater than result
			if len(result) > len(str[l:i+1]) {
				result = str[l : i+1]
			}

			// slide window to next character
			lch := rune(str[l])
			strMap[lch]--
			diff[lch]++
			l++
		}
	}
	return result
}

func sqeezeFromLeft(str string, charMap, strMap map[rune]int) int {
	squeezeBy := 0
	for _, ch := range str {
		if val, ok := charMap[ch]; ok {
			if strMap[ch] == val {
				break
			}
			strMap[ch]--
		}
		squeezeBy++
	}
	return squeezeBy
}

func createMap(str string) map[rune]int {
	strMap := make(map[rune]int)
	for _, v := range str {
		strMap[v]++
	}
	return strMap
}
