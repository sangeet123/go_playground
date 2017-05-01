package rotatearr

func Rotate(arr []int, rightShift int) {

	if rightShift < 0 {
		panic("usuage: value of rightShift cannot be negative")
	}

	if len(arr) < 2 {
		return
	}

	actualShift := rightShift % len(arr)

	if rightShift == 0 {
		return
	}

	rotate(arr[:])
	rotate(arr[:actualShift])
	rotate(arr[actualShift:])
}

func rotate(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
