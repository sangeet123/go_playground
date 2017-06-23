package tree

func (f FullTree) LeastCommonAncestor(a, b int) (int, bool) {
	afound := false
	bfound := false
	common := leastCommonAncestor(a, b, &afound, &bfound, f.root)
	if afound && bfound {
		return common.data, true
	}
	return -1, false
}

func leastCommonAncestor(a, b int, afound, bfound *bool, n *node) *node {
	if n == nil {
		return nil
	}
	l := leastCommonAncestor(a, b, afound, bfound, n.l)
	r := leastCommonAncestor(a, b, afound, bfound, n.r)

	if l != nil && r != nil {
		return n
	}

	if *afound && *bfound {
		return returnNonNullOne(l, r)
	}

	if !*afound && a == n.data {
		*afound = true
		return n
	}

	if !*bfound && b == n.data {
		*bfound = true
		return n
	}
	return returnNonNullOne(l, r)
}

func returnNonNullOne(l, r *node) *node {
	if l != nil {
		return l
	}
	return r
}
