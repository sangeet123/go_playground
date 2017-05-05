package sort

type integers []int

func BubbleSortInt(nos integers) {
	for i := 0; i < len(nos)-1; i++ {
		for j := 0; j < len(nos)-i-1; j++ {
			if nos[j] > nos[j+1] {
				nos[j], nos[j+1] = nos[j+1], nos[j]
			}
		}
	}
}
