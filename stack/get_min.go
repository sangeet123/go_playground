package stack

type minstack struct {
	s   *stack
	min *stack
}

func NewMinStack(capacity int) *minstack {
	minstack := new(minstack)
	minstack.s = NewStack(capacity)
	minstack.min = NewStack(capacity)
	return minstack
}

func (this minstack) Size() {
	this.s.Size()
}

// Only allowing to push integer
// Can be extended to any type
func (this minstack) Push(data int) {
	this.s.Push(data)
	dataToPush := data
	if !this.min.IsEmpty() {
		topData := this.min.Peek().(int)
		if topData < data {
			dataToPush = topData
		}
	}
	this.min.Push(dataToPush)
}

func (this minstack) Pop() int {
	this.min.Pop()
	return this.s.Pop().(int)
}

func (this minstack) GetMinInt() int {
	return this.min.Peek().(int)
}

func (this minstack) IsEmpty() bool {
	return this.s.IsEmpty()
}

func (this minstack) IsFull() bool {
	return this.s.IsFull()
}
