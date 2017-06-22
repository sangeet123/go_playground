package arrayprob

// RotateRightBy function to rotate array
func RotateRightBy(arr []int, k int) {

	if k < 0 {
		panic("value of k cannot be negative")
	}
	if len(arr) == 0 {
		return
	}
	if k = k % len(arr); k == 0 {
		return
	}
	Reverse(arr)
	Reverse(arr[0:k])
	Reverse(arr[k:])
}

//function that reverses an array
func Reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
