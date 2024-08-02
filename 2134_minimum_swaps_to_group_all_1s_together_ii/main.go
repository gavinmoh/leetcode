package main

func minSwaps(nums []int) int {
	ones := 0
	for _, num := range nums {
		if num == 1 {
			ones++
		}
	}

	count := 0
	for i := 0; i < ones; i++ {
		if nums[i] == 0 {
			count++
		}
	}

	currentCount := count
	nums = append(nums, nums...)
	for i := 1; i < len(nums)-ones; i++ {
		if nums[i-1] == 0 {
			currentCount--
		}
		if nums[i+ones-1] == 0 {
			currentCount++
		}
		count = min(currentCount, count)
	}

	return count
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
