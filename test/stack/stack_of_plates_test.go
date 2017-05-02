package stackofplatestest

import (
	"go_playground/stack"
	"testing"
)

func TestGetMinIntOp(t *testing.T) {
	s := stack.NewStackOfPlates(2)
	valuesInStack := []int{3, 4, 5, 1, 0, 2, -1, -2, -3, 4}

	for _, val := range valuesInStack {
		s.Push(val)
	}

	for i := 9; i >= 0; i-- {
		val := s.Pop()
		if val != valuesInStack[i] {
			t.Error("Expected i:", i, valuesInStack[i], " But got ", val)
		}
	}
}
