package queuetest

import (
	"go_playground/stack"
	"testing"
)

func TestQueueOperationsEnqueFollowedByDequeue(t *testing.T) {
	q := stack.NewQueue(10)

	for i := 0; i < 10; i++ {
		q.Enqueue(i)
	}

	for i := 0; i < 10; i++ {
		data := q.Dequeue()

		if data != i {
			t.Error("s.Dequeue(): Expected ", i, " but got ", data)
		}
	}
}

func TestQueueRandomOperations(t *testing.T) {
	q := stack.NewQueue(10)

	q.Enqueue(0)
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)

	val := q.Dequeue()
	if val != 0 {
		t.Error("s.Dequeue(): Expected ", 0, " but got ", val)
	}

	q.Enqueue(5)

	val = q.Dequeue()

	if val != 1 {
		t.Error("s.Dequeue(): Expected ", 1, " but got ", val)
	}

	q.Enqueue(6)
	q.Enqueue(7)

	val = q.Dequeue()
	if val != 2 {
		t.Error("s.Dequeue(): Expected ", 2, " but got ", val)
	}

	val = q.Dequeue()
	if val != 3 {
		t.Error("s.Dequeue(): Expected ", 3, " but got ", val)
	}

	for i := 4; i < 8; i++ {
		data := q.Dequeue()

		if data != i {
			t.Error("s.Dequeue(): Expected ", i, " but got ", data)
		}
	}
}
