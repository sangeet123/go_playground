package tree

func (f FullTree) GetPath(sum int) [][]int {
	result := [][]int{}
	getPath(f.root, []int{0}, []int{0}, &result, sum)
	return result
}

func getPath(n *node, path []int, sumPath []int, result *[][]int, sum int) {
	if n == nil {
		return
	}
	path = append(path, n.data)
	sumPath = append(sumPath, sumPath[len(sumPath)-1]+n.data)
	process(path, sumPath, result, sum)
	getPath(n.l, path, sumPath, result, sum)
	getPath(n.r, path, sumPath, result, sum)
}

func process(path []int, sumPath []int, result *[][]int, sum int) {
	j := len(path) - 1
	for i := 1; i < len(path); i++ {
		sumij := sumPath[j] - sumPath[i-1]
		if sumij == sum {
			toAppend := []int{}
			for ii := i; ii <= j; ii++ {
				toAppend = append(toAppend, path[ii])
			}
			*result = append(*result, toAppend)
		}
	}
}
