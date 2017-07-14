package tree

func (f *FullTree) Invert() {
	invert(f.root)
}

func invert(n *node) {
	if n == nil {
		return
	}
	temp := n.l
	n.l = n.r
	n.r = temp
	invert(n.l)
	invert(n.r)
}
