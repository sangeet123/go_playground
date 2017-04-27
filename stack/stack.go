package stack

type custom_type interface{}

type StackInterface interface {
	Push(data custom_type)
	Pop() custom_type
	Peek() custom_type
	Size() int
	IsEmpty() bool
}

func NewStack() *stack {
	s := new(stack)
	s.top = -1
	return s
}

type stack struct {
	data []custom_type
	top  int
	min  *stack
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

func (this *stack) Push(data custom_type) {
	this.data = append(this.data, data)
	this.top++
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
