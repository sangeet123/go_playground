package sort

type integersPtr *[]int

func MergeSortInt(nos integers) {
	mergeSortInt(0, len(nos)-1, nos)
}

func mergeSortInt(l int, r int, nos integers) {
	if l < r {
		m := (l + r) >> 1
		mergeSortInt(l, m, nos)
		mergeSortInt(m+1, r, nos)

		temp := []int{}

		i := l
		j := m + 1

		for i <= m && j <= r {
			if nos[i] <= nos[j] {
				temp = append(temp, nos[i])
				i++
			} else {
				temp = append(temp, nos[j])
				j++
			}
		}
		assignRem(i, m, nos, &temp)
		assignRem(j, r, nos, &temp)
		for k, i := 0, l; i <= r; i, k = i+1, k+1 {
			nos[i] = temp[k]
		}

	}
}

func assignRem(fromInd int, toInd int, from integers, to integersPtr) {
	for i := fromInd; i <= toInd; i++ {
		*to = append(*to, from[i])
	}
}
