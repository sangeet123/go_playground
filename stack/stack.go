package stack

type custom_type interface{}

type StackInterface interface {
	Push(data custom_type)
	Pop() custom_type
	Peek() custom_type
	Size() int
	IsEmpty() bool
}

// There is no reallocation once size is fixed
// stack wont grow pas the capacity of stack
func NewStack(capacity int) *stack {
	s := new(stack)
	s.data = make([]custom_type, 0, capacity)
	return s
}

type stack struct {
	data []custom_type
}

func (this stack) Peek() custom_type {
	if this.Size() == 0 {
		panic("stack is empty")
	}

	return this.data[this.Size()-1]
}

func (this stack) Size() int {
	return len(this.data)
}

func (this stack) IsEmpty() bool {
	return len(this.data) == 0
}

func (this stack) IsFull() bool {
	return this.Size() == cap(this.data)
}

func (this *stack) Push(data custom_type) {
	if this.IsFull() {
		panic("stack is full")
	}
	this.data = append(this.data, data)
}

func (this *stack) Pop() custom_type {
	if this.IsEmpty() {
		panic("stack is empty")
	}
	data := this.data[this.Size()-1]
	this.data = this.data[:this.Size()-1]
	return data
}
