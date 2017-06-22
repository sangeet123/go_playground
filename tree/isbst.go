package tree

import (
	"math"
)

func (f FullTree) IsBST() bool {
	return isBST(f.root, -math.MaxInt64, math.MaxInt64)
}

func isBST(n *node, min, max int) bool {
	if n == nil {
		return true
	}
	return min < n.data && n.data <= max && (n == nil || isBST(n.l, min, n.data) && isBST(n.r, n.data, max))
}
