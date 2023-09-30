package main

func find132pattern(nums []int) bool {
	stack := []int{}
	candidateK := -(1 << 31)

	for i := len(nums) - 1; i >= 0; i-- {

		// if nums[i] (representing j) is less than candidateK (representing k),
		// then we found a valid i, j, k
		if nums[i] < candidateK {
			return true
		}

		// if nums[i] (representing j) is greater than the top of the stack,
		// then we can pop the stack and update candidateK
		for len(stack) > 0 && nums[i] > stack[len(stack)-1] {
			candidateK = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, nums[i])

	}

	return false
}
