package sort

func InsertionSortInt(nos integers) {
	for i := 1; i < len(nos); i++ {
		curNo := nos[i]
		j := i - 1
		for ; j >= 0 && curNo < nos[j]; j-- {
			nos[j+1] = nos[j]
		}
		nos[j+1] = curNo
	}
}
