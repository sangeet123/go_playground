package arrayprob

import (
	"go_playground/stack"
)

func LongestValidParen(str string) string {
	l := len(str)
	s := stack.NewStack(l)
	arr := []int{}
	for i := 0; i < l; i++ {
		val := rune(str[i])
		if IsLeftParen(val) {
			s.Push(i)
		} else if (IsRightParen(val)) && !s.IsEmpty() {
			ind := s.Pop().(int)
			lParen := rune(str[ind])
			if AreValidPair(lParen, val) {
				arr = append(arr, []int{ind, i}...)
			} else {
				s.Clear()
			}
		} else {
			s.Clear()
		}
	}

	if len(arr) == 0 {
		return ""
	}
	start, end := GetLargestRange(arr)
	return str[start : end+1]
}

// Uses counting sort to find the largest range as we know
// the range of the element cannot exceed string length
func GetLargestRange(arr []int) (int, int) {
	sorted := make([]bool, arr[len(arr)-1]+1)
	for i := 0; i < len(arr); i++ {
		sorted[arr[i]] = true
	}
	ra, ca, rb, cb := 0, 0, 0, 0
	for i := 1; i < len(sorted); i++ {
		if sorted[i] {
			cb = i
		} else {
			if cb-ca > rb-ra {
				ra, rb = ca, cb
			}
			ca = i + 1
		}
	}
	if cb-ca > rb-ra {
		return ca, cb
	}
	return ra, rb
}

func IsLeftParen(paren rune) bool {
	return paren == '(' || paren == '[' || paren == '{'
}

func IsRightParen(paren rune) bool {
	return paren == ')' || paren == ']' || paren == '}'
}

func AreValidPair(lparen, rparen rune) bool {
	return lparen == '(' && rparen == ')' || lparen == '{' && rparen == '}' || lparen == '[' && rparen == ']'
}
