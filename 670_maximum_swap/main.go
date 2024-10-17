package main

func maximumSwap(num int) int {
	nums := []int{}
	for num > 0 {
		nums = append(nums, num%10)
		num /= 10
	}

	max := make([]int, len(nums))
	maxIndex := -1
	maxIndices := make([]int, len(nums))
	currentMax := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] > currentMax {
			currentMax = nums[i]
			maxIndex = i
		}
		max[i] = currentMax
		maxIndices[i] = maxIndex
	}

	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] < max[i] {
			// swap
			nums[i], nums[maxIndices[i]] = nums[maxIndices[i]], nums[i]
			break
		}
	}

	for i := len(nums) - 1; i >= 0; i-- {
		num = (num * 10) + nums[i]
	}

	return num
}
