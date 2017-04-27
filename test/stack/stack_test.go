package stacktest

import (
	"go_playground/stack"
	"testing"
)


func TestStackOperations(t *testing.T) {
	s := stack.NewStack()

	for i := 0 ; i < 10 ; i++{
		s.Push(i)
	}

	for i := 9 ; i >= 0 ; i--{
		data := s.Pop()

		if data != i{
			t.Error("Expected ", i , " but got ", data)
		}
	}

	if s.Size() != 0 {
		t.Error("Expected size of 0 but got ", s.Size())
	}

	poping from empty stack should result in error
	defer func() {
		if r := recover(); r == nil {
			t.Error("no error ocurred")
		}
	}()

	s.Pop()

}