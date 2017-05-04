package stack

func (this stack) SortInt() {
	if this.Size() <= 1 {
		return
	}

	temp := NewStack(this.Size())

	for !this.IsEmpty() {
		thisTop := this.Pop().(int)
		for !temp.IsEmpty() {
			tempTop := temp.Peek().(int)
			if tempTop < thisTop {
				break
			}
			this.Push(tempTop)
			temp.Pop()
		}
		temp.Push(thisTop)
	}

	for !temp.IsEmpty() {
		this.Push(temp.Pop())
	}

}
