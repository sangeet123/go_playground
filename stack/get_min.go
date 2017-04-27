package stack

func (this stack) GetMinInt() int {
	if this.Size() == 0 {
		panic("Stack is empty")
	}

	if this.min == nil {
		this.min = NewStack()
		this.min.Push(this.data[0])
	}

	if this.min.Size() > this.Size() {
		this.min.data = this.min.data[:this.top+1]
		this.min.top = this.top
	}

	if this.min.Size() < this.Size() {
		for i := this.min.top; i < this.Size(); i++ {
			min, ok := this.min.Peek().(int)
			if !ok {
				this.min = nil
				panic("Stack has items other than integer")
			}

			if min < this.data[i].(int) {
				this.min.Push(min)
			} else {
				this.min.Push(this.data[i])
			}
		}
	}

	return this.min.Peek().(int)
}
