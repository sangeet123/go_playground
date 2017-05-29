package tree

// return new node
func getNode(data int) *node {
	n := new(node)
	n.data = data
	n.l = nil
	n.r = nil
	n.p = nil
	return n
}

func linkWithParent(p, c *node) {
	if c != nil {
		c.p = p
	}
}

func makeLeftChild(p, c *node) *node {
	if p != nil {
		p.l = c
		linkWithParent(p, c)
		return p
	}
	return c
}

func makeRightChild(p, c *node) *node {
	if p != nil {
		p.r = c
		linkWithParent(p, c)
		return p
	}
	return c
}

func (n node) getInorderPredecessor() *node {
	start := n.r
	for start.l != nil {
		start = start.l
	}
	return start
}

func (n node) isRootNode() bool {
	return n.p == nil
}

func (n node) isLeftChild() bool {
	return n.p != nil && n.p.l == &n
}

func (n node) isRightChild() bool {
	return n.p != nil && n.p.r == &n
}

func (n node) hasLeftChild() bool {
	return n.l != nil
}

func (n node) hasRightChild() bool {
	return n.r != nil
}

func (n node) hasBothChild() bool {
	return n.hasLeftChild() && n.hasRightChild()
}

func (n node) isLeafNode() bool {
	return !n.hasLeftChild() || !n.hasRightChild()
}

func (n *node) stripAllLinks() {
	n.p = nil
	n.l = nil
	n.r = nil
}
