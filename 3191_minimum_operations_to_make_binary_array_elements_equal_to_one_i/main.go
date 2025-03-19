package main

func minOperations(nums []int) int {
	n := len(nums)
	count := 0

	flip := func(i int) {
		if nums[i] == 0 {
			nums[i] = 1
		} else {
			nums[i] = 0
		}
	}

	for i := 0; i < n-2; i++ {
		if nums[i] == 1 {
			continue
		}

		count++
		flip(i)
		flip(i + 1)
		flip(i + 2)
	}

	if nums[n-1] == 0 || nums[n-2] == 0 {
		return -1
	}

	return count
}
