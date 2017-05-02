package stack

type queue struct {
	s1 *stack
	s2 *stack
}

//user would be seeing the same error message of
// stack full and empty
func NewQueue(capacity int) *queue {
	q := new(queue)
	q.s1 = NewStack(capacity)
	q.s2 = NewStack(capacity)
	return q
}

func (this *queue) Enqueue(data custom_type) {
	this.s1.Push(data)
}

func (this *queue) Dequeue() custom_type {
	if this.s2.IsEmpty() {
		for !this.s1.IsEmpty() {
			data := this.s1.Pop()
			this.s2.Push(data)
		}
	}
	return this.s2.Pop()
}
