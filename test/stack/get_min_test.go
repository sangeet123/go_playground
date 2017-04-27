package getmintest

import (
	"go_playground/stack"
	"testing"
)

func TestGetMinIntOp(t *testing.T) {
	s := stack.NewStack()
	valuesInStack := []int{3, 4, 5, 1, 0, 2, -1, -2, -3, 4}
	minValues := []int{3, 3, 3, 1, 0, 0, -1, -2, -3, -3}

	for _, val := range valuesInStack {
		s.Push(val)
	}

	i := 9
	for i >= 5 {
		if s.GetMinInt() != minValues[i] {
			t.Error("Expected ", minValues[i], " But got ", s.GetMinInt())
		}
		s.Pop()
		i--
	}

	s.Pop()
	s.Pop()
	i--
	i--

	for i >= 0 {
		if s.GetMinInt() != minValues[i] {
			t.Error("Expected ", minValues[i], " But got ", s.GetMinInt())
		}
		s.Pop()
		i--
	}
}
