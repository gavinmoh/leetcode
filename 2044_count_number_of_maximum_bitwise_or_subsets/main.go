package main

func countMaxOrSubsets(nums []int) int {
	var helper func(index int, current int, target int) int
	helper = func(index int, current int, target int) int {
		if index == len(nums) {
			if current == target {
				return 1
			}
			return 0
		}

		left := helper(index+1, current, target)
		right := helper(index+1, current|nums[index], target)

		return left + right
	}

	max := 0
	for _, num := range nums {
		max |= num
	}

	return helper(0, 0, max)
}
