package rotated

import "strings"

func is_rotated_of(word string, rotated string)bool{
	if len(word) != len(rotated){
		return false
	}
	concatenated := word + rotated
	return strings.Contains(concatenated, rotated)
}