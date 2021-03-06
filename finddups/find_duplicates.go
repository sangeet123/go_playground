package finddups

type Integers []int

func GetFirstDup(nums Integers) int {
	for _, val := range nums {
		ind := abs(val) - 1

		if nums[ind] < 0 {
			return ind + 1
		}

		nums[ind] = -nums[ind]
	}
	return 0
}

func abs(val int) int {
	if val >= 0 {
		return val
	}
	return -val
}
