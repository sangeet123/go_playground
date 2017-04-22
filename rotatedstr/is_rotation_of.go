package rotatedstr

import "strings"

func IsRotated(word string, rotated string) bool {
	if len(word) != len(rotated) {
		return false
	}
	concatenated := word + rotated
	return strings.Contains(concatenated, rotated)
}
