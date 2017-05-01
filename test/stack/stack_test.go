package stacktest

import (
	"go_playground/stack"
	"testing"
)

func TestStackOperations(t *testing.T) {
	s := stack.NewStack(10)

	for i := 0; i < 10; i++ {
		s.Push(i)
	}

	if s.IsFull() != true {
		t.Error("s.IsFull(): Expected true but received false")
	}

	for i := 9; i >= 0; i-- {
		data := s.Peek()
		s.Pop()

		if data != i {
			t.Error("s.Pop(): Expected ", i, " but got ", data)
		}
	}

	if s.Size() != 0 {
		t.Error("s.Size():Expected size of 0 but got ", s.Size())
	}

	if s.IsEmpty() != true {
		t.Error("s.IsEmpty():Expected true but got false")
	}

	//poping from empty stack should result in error
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic but no error ocurred")
		}
	}()

	s.Pop()

}

func TestOperationsForEmptyStackOperations(t *testing.T) {
	s := stack.NewStack(0)

	if s.Size() != 0 {
		t.Error("Expected size of 0 but got ", s.Size())
	}

	if s.IsEmpty() != true {
		t.Error("Expected true but got false")
	}

	//poping from empty stack should result in error
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic but no error ocurred")
		}
	}()

	s.Pop()

}

func TestPushOnFullStack(t *testing.T) {
	s := stack.NewStack(1)
	//poping from empty stack should result in error
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic but no error ocurred")
		}
	}()

	s.Push(1)
	s.Push(2)

}
