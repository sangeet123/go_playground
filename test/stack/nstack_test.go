package nstacktest

import (
	"go_playground/stack"
	"testing"
)

func TestStackOperations(t *testing.T) {
	s := stack.NewNStack(4, 20)

	for i := 0; i < 2; i++ {
		s.Push(0, 12)
		s.Push(1, 13)
		s.Push(2, 14)
		s.Push(3, 15)

		s.Push(0, 16)
		s.Push(1, 17)
		s.Push(2, 18)
		s.Push(3, 19)

		val := s.Pop(0)

		if val != 16 {
			t.Error("Expected 16 but got ", val)
		}

		val = s.Pop(1)
		if val != 17 {
			t.Error("Expected 17 but got ", val)
		}

		val = s.Pop(2)
		if val != 18 {
			t.Error("Expected 18 but got ", val)
		}

		val = s.Pop(3)
		if val != 19 {
			t.Error("Expected 19 but got ", val)
		}

		val = s.Pop(0)

		if val != 12 {
			t.Error("Expected 12 but got ", val)
		}

		val = s.Pop(1)
		if val != 13 {
			t.Error("Expected 13 but got ", val)
		}

		val = s.Pop(2)
		if val != 14 {
			t.Error("Expected 14 but got ", val)
		}

		val = s.Pop(3)
		if val != 15 {
			t.Error("Expected 15 but got ", val)
		}
	}

}

func TestPopOnEmptyStack(t *testing.T) {
	s := stack.NewNStack(4, 20)

	defer func() {
		if r := recover(); r == nil {
			t.Error("no error ocurred")
		}
	}()

	s.Pop(1)
}

func TestPushOnStackThatDoesNotExist(t *testing.T) {
	s := stack.NewNStack(4, 20)

	defer func() {
		if r := recover(); r == nil {
			t.Error("no error ocurred")
		}
	}()

	s.Pop(4)
}

func TestPushAndPopOperations(t *testing.T) {
	s := stack.NewNStack(4, 2)

	s.Push(0, 1)
	s.Push(0, 2)

	val := s.Pop(0)
	if val != 2 {
		t.Error("Expected 1 but got ", val)
	}

	val = s.Pop(0)
	if val != 1 {
		t.Error("Expected 2 but got ", val)
	}

	s.Push(1, 1)
	s.Push(2, 1)

	defer func() {
		if r := recover(); r == nil {
			t.Error("no error ocurred")
		}
	}()

	s.Push(3, 2)

}
