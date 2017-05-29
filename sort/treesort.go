package sort

import "go_playground/tree"

func comp(a, b int) int {
	if a == b {
		return 0
	} else if a > b {
		return 1
	}
	return -1
}

func TreeSortInt(nos integers) {
	intTree := tree.GetIntTree(comp)
	for i := 0; i < len(nos); i++ {
		intTree.Insert(nos[i])
	}

	i := 0
	for _, val := range intTree.InOrder() {
		nos[i] = val
		i++
	}
}
