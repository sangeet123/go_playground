package tree

func IsSubTree(f1, f2 *FullTree) bool {

	if f1.root == nil && f2.root == nil {
		return true
	}

	if f1.root == nil || f2.root == nil {
		return false
	}

	return isSubTree(f1.root, f2.root)
}

func isSubTree(n1 *node, n2 *node) bool {
	if n1 != nil {
		isValid := false
		if n1.data == n2.data {
			isValid = processDescendents(n1, n2)
		}
		return isValid || isSubTree(n1.l, n2) || isSubTree(n1.r, n2)
	}
	return false
}

func processDescendents(n1 *node, n2 *node) bool {
	return (n1 == nil && n2 == nil) ||
		(n1 != nil && n2 != nil &&
			n1.data == n2.data &&
			processDescendents(n1.l, n2.l) &&
			processDescendents(n1.r, n2.r))
}
