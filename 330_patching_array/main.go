package main

func minPatches(nums []int, n int) int {
	i, count, max := 0, 0, 0
	for max < n {
		if i < len(nums) && nums[i] <= max+1 {
			max += nums[i]
			i++
		} else {
			count++
			max += max + 1
		}

	}

	return count
}
