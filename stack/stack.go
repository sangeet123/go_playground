package stack

type custom_type interface{}

type StackInterface interface {
	Push(data custom_type)
	Pop() custom_type
	Peek() custom_type
	Size() int
	IsEmpty() bool
}

func NewStack(capacity int) *stack {
	s := new(stack)
	s.top = -1
	s.data = make([]custom_type, capacity, capacity)
	return s
}

type stack struct {
	data []custom_type
	top  int
}

func (this stack) Peek() custom_type {
	if this.Size() == 0 {
		panic("stack is empty")
	}

	return this.data[this.top]
}

func (this stack) Size() int {
	return this.top + 1
}

func (this stack) IsEmpty() bool {
	return this.top == -1
}

func (this stack) IsFull() bool {
	return this.Size() == cap(this.data)
}

func (this *stack) Push(data custom_type) {
	if this.IsFull() {
		panic("stack is full")
	}
	this.top++
	this.data[this.top] = data
}

func (this *stack) Pop() custom_type {
	if this.top == -1 {
		panic("stack is empty")
	}
	data := this.data[this.top]
	this.data = this.data[:this.top]
	this.top--
	return data
}
