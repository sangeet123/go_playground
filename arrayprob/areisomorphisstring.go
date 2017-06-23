package arrayprob

func AreIsomorphicStrings(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	strMap := make(map[byte]byte)

	for i := 0; i < len(str1); i++ {
		if val, ok := strMap[str1[i]]; ok && val != str2[i] {
			return false
		}

		if val, ok := strMap[str2[i]]; ok && val != str1[i] {
			return false
		}
		strMap[str1[i]] = str2[i]
		strMap[str2[i]] = str1[i]
	}

	return true
}
