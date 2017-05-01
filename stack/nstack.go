package stack

type NStack struct {
	data              []custom_type
	top_of_stack      [][]int
	next_free         []int
	next_free_pointer int
	no_of_stack       int
}

func NewNStack(n, size int) *NStack {
	nStack := new(NStack)
	nStack.no_of_stack = n
	nStack.data = make([]custom_type, size, size)
	nStack.top_of_stack = make([][]int, n)
	nStack.next_free = make([]int, size)
	for i := 0; i < size-1; i++ {
		nStack.next_free[i] = i + 1
	}
	nStack.next_free[size-1] = -1
	return nStack
}

func (this NStack) assertValidStack(i int) {
	if i < 0 && i >= this.no_of_stack {
		panic("not a valid stack")
	}
}

func (this NStack) Size(i int) int {
	this.assertValidStack(i)
	return len(this.top_of_stack[i])
}

func (this *NStack) Push(i int, data custom_type) {
	this.assertValidStack(i)
	free_loc := this.next_free_pointer

	if free_loc == -1 {
		panic("Stack is full")
	}

	this.next_free_pointer = this.next_free[free_loc]
	this.top_of_stack[i] = append(this.top_of_stack[i], free_loc)
	this.data[free_loc] = data
}

func (this *NStack) Pop(i int) custom_type {
	this.assertValidStack(i)
	len := this.Size(i)
	if len == 0 {
		panic("Stack is empty")
	}

	top := this.top_of_stack[i][len-1]
	to_return := this.data[top]
	this.next_free[top] = this.next_free_pointer
	this.next_free_pointer = top
	this.top_of_stack[i] = this.top_of_stack[i][:len-1]
	return to_return
}
