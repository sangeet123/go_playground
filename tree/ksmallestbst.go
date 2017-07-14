package tree

func GetSmallest(this *intTree, k int) int {
	if k <= 0 {
		panic("k must be greater than 0")
	}

	index, kth := k, 0
	getSmallest(this.root, &index, &kth)

	if index > 0 {
		panic("k must be  within total elements in array")
	}

	return kth
}

func getSmallest(n *node, index, kth *int) {
	if n == nil || *index <= 0 {
		return
	}
	getSmallest(n.l, index, kth)
	*index = *index - 1
	if *index == 0 {
		*kth = n.data
	}
	getSmallest(n.r, index, kth)
}
