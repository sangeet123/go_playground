package tree

// sorted in ascending order
func assertSorted(arr []int) {
	for j := 1; j < len(arr); j++ {
		if arr[j] < arr[j-1] {
			panic("not a sorted array")
		}
	}
}

func comparator(a, b int) int {
	if a == b {
		return 0
	} else if a > b {
		return 1
	}
	return -1
}

func CreateMinimumTree(arr []int) *intTree {
	assertSorted(arr)
	intTree := GetIntTree(comparator)
	intTree.root = createTree(arr, 0, len(arr)-1)
	return intTree
}

func createTree(arr []int, l int, r int) *node {
	if l <= r {
		m := (l + r) >> 1
		newNode := getNode(arr[m])
		newNode.l = createTree(arr, l, m-1)
		if newNode.l != nil {
			newNode.l.p = newNode
		}
		newNode.r = createTree(arr, m+1, r)
		if newNode.r != nil {
			newNode.r.p = newNode
		}
		return newNode
	}
	return nil
}
