package sorttest

import (
	"go_playground/stack"
	"testing"
)

func TestSortStack(t *testing.T) {
	s := stack.NewStack(10)

	for i := 1; i <= 10; i++ {
		s.Push(i)
	}

	s.SortInt()

	for i := 1; i <= 10; i++ {
		data := s.Peek()
		s.Pop()
		if data != i {
			t.Error("s.Pop(): Expected ", i, " but got ", data)
		}
	}

}

func TestSortStackWithSingleElement(t *testing.T) {
	s := stack.NewStack(1)

	s.Push(1)

	s.SortInt()

	if s.Peek() != 1 {
		t.Error("Expected 1 but got ", s.Peek())
	}
}
