package multiplyallbutself

type Integers []int

func (slice Integers) Len() int {
	return len(slice)
}

func MultiplyAllButSelf(nums Integers) Integers {

	length := nums.Len()

	// if length is empty return empty array
	if length == 0 {
		return nil
	}

	// if length is one return 0
	if length == 1 {
		return Integers{0}
	}

	result := make(Integers, length)
	result[0] = 1

	for i := 1; i < length; i++ {
		result[i] = nums[i-1] * result[i-1]
	}

	back := 1
	for i := length - 1; i >= 0; i-- {
		result[i] = result[i] * back
		back = back * nums[i]
	}

	return result
}
