package tree

// InOrderSuccessor returns if tree is balance or not
func (this *intTree) InOrderSuccessor(data int) (int, bool) {
	_, n := this.find(data)

	// no such node found
	if n == nil {
		return -1, false
	}

	// if it is not a left child, see if it has
	// right child and go left till null and return that node
	if n.hasRightChild() {
		r := n.r
		for r.l != nil {
			r = r.l
		}
		return r.data, true
	}

	// it it has no right child go up until
	// node that is a left child of a parent comes
	for n != nil && n.isRightChild() {
		n = n.p
	}

	if n != nil && n.p != nil {
		return n.p.data, true
	}

	return -1, false
}
