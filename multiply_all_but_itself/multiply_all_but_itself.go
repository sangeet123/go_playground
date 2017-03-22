package multiplyallbutitself

func multiply_all_but_itself(nums []int)[]int{
	
	length := len(nums)

	// if length is empty return empty array
	if length == 0{
		return []int{}
	}

	// if length is one return 0
	if length == 1{
		return []int{0}
	}

	result := make([]int, length)
	result[0] = 1

	for i:=1; i < length; i++{
		result[i] = nums[i-1] * result[i-1]
	}

	back := 1
	for i := length - 1 ; i >= 0; i--{
		result[i] = result[i] * back
		back = back * nums[i]
	}

	return result
}