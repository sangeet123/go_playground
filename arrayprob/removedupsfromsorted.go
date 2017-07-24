package arrayprob

// Removes duplicate from array that allows maximum of allow element
func RemoveDups(arr []int, allow int) []int {
	if allow <= 0 || len(arr) == 0 {
		return nil
	}

	count := 1
	cur := arr[0]
	l := 1

	for i := 1; i < len(arr); i++ {
		if arr[i] != cur {
			cur = arr[i]
			count = 1
			arr[l] = arr[i]
			l++
		} else if arr[i] == cur {
			if count < allow {
				count++
				arr[l] = arr[i]
				l++
			} else {
				count++
			}
		}
	}

	return arr[0:l]
}
