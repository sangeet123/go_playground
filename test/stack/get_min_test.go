package getmintest

import (
	"go_playground/stack"
	"testing"
)

func TestGetMinIntOp(t *testing.T) {
	s := stack.NewMinStack(10)
	valuesInStack := []int{3, 4, 5, 1, 0, 2, -1, -2, -3, 4}
	minValues := []int{3, 3, 3, 1, 0, 0, -1, -2, -3, -3}

	for _, val := range valuesInStack {
		s.Push(val)
	}

	for i := 9; i >= 0; i-- {
		if s.GetMinInt() != minValues[i] {
			t.Error("Expected ", minValues[i], " But got ", s.GetMinInt())
		}
		s.Pop()
	}
}
