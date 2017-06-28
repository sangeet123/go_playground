package arrayprob

func MatchesRegex(word, pattern string) bool {
	return matchesRegex(word, pattern, 0, 0)
}

func matchesRegex(word, pattern string, wind, pind int) bool {
	if wind >= len(word) && pind >= len(pattern) {
		return true
	}

	// we are already pass word and the final value is asterisk we
	// we have matched the pattern
	if wind >= len(word) {
		if pind == len(pattern)-1 && pattern[pind] == '*' {
			return true
		}
		return false
	}

	if pind >= len(pattern) {
		return false
	}

	if pattern[pind] == '.' || pattern[pind] == word[wind] {
		return matchesRegex(word, pattern, wind+1, pind+1)
	} else if pattern[pind] == '*' {
		return matchesRegex(word, pattern, wind, pind+1) || matchesRegex(word, pattern, wind, pind-1)
	} else if pind+1 < len(pattern) && pattern[pind+1] == '*' {
		return matchesRegex(word, pattern, wind, pind+2)
	}
	return false
}
