package highestproduceofthreeintegers

import "sort"

type Integers []int

func (slice Integers) Len() int {
	return len(slice)
}

func get_highest_product(nums Integers) int {

	if nums.Len() < 3 {
		panic("no of elements should be greater than equal to 3")
	}

	if nums.Len() == 3 {
		return nums[0] * nums[1] * nums[2]
	}

	// The logic seems quite complicated but it is not
	// if all +ve numbers the three largest elements gives the largest product
	// if there are negatives find three largest as well as two smallest
	// return the max of (largest * product of two smallest, product of three largest)
	first_three := nums[:3]
	sort.Ints(first_three)
	max_1, max_2, max_3 := first_three[2], first_three[1], first_three[0]
	min_1, min_2 := first_three[0], first_three[1]
	for i := 3; i < nums.Len(); i++ {
		// maintain three largest
		if max_1 < nums[i] {
			max_1, max_2, max_3 = nums[i], max_1, max_2
		} else if max_2 < nums[i] {
			max_2, max_3 = nums[i], max_2
		} else if max_3 < nums[i] {
			max_3 = nums[i]
		}

		// maintain two smallest
		if min_1 > nums[i] {
			min_1, min_2 = nums[i], min_1
		} else if min_2 > nums[i] {
			min_2 = nums[i]
		}
	}

	productmax123 := max_1 * max_2 * max_3
	productmax12min1 := max_1 * min_1 * min_2

	return max(productmax123, productmax12min1)
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
