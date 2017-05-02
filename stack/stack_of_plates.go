package stack

type stacks []*stack

type stackOfPlates struct {
	s                stacks
	capacityPerStack int
	curStackInd      int
}

func NewStackOfPlates(capacity int) *stackOfPlates {
	stcs := new(stackOfPlates)
	stcs.s = append(stcs.s, NewStack(capacity))
	stcs.capacityPerStack = capacity
	stcs.curStackInd = 0
	return stcs
}

func (this *stackOfPlates) Push(data custom_type) {
	i := this.curStackInd
	if cap(this.s[i].data) == this.s[i].Size() {
		this.s = append(this.s, NewStack(this.capacityPerStack))
		this.curStackInd, i = i+1, i+1
	}
	this.s[i].Push(data)
}

func (this *stackOfPlates) Pop() custom_type {
	i := this.curStackInd
	if this.s[i].Size() == 0 {
		if i == 0 {
			panic("no stack on a stack")
		}
		this.s = this.s[:this.curStackInd]
		this.curStackInd, i = i-1, i-1
	}
	return this.s[i].Pop()
}
