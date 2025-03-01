package main

func applyOperations(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			nums[i] *= 2
			nums[i+1] = 0
		}
	}

	for left := 0; left < len(nums)-1; left++ {
		if nums[left] != 0 {
			continue
		}

		for right := left + 1; right < len(nums); right++ {
			if nums[right] != 0 {
				nums[left], nums[right] = nums[right], nums[left]
				break
			}
		}
	}

	return nums
}
