package main

func countGood(nums []int, k int) int64 {
	n := len(nums)
	pair, right := 0, -1
	freq := map[int]int{}
	count := int64(0)
	for left := 0; left < n; left++ {
		for pair < k && right+1 < n {
			right++
			pair += freq[nums[right]]
			freq[nums[right]]++
		}
		if pair >= k {
			count += int64(n - right)
		}
		freq[nums[left]]--
		pair -= freq[nums[left]]
	}

	return count
}
