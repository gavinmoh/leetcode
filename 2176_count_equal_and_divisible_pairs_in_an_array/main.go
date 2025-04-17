package main

func countPairs(nums []int, k int) int {
	n := len(nums)
	count := 0

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i] == nums[j] && (i*j)%k == 0 {
				count++
			}
		}
	}

	return count
}
