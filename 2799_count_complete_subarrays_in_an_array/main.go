package main

func countCompleteSubarrays(nums []int) int {
	distinct := map[int]bool{}
	for _, num := range nums {
		distinct[num] = true
	}
	distinctCount := len(distinct)

	result := 0
	freq := map[int]int{}
	n, right := len(nums), 0
	for left := 0; left < n; left++ {
		if left > 0 {
			freq[nums[left-1]]--
			if freq[nums[left-1]] == 0 {
				delete(freq, nums[left-1])
			}
		}

		// expand the window
		for right < n && len(freq) < distinctCount {
			freq[nums[right]]++
			right++
		}

		if len(freq) == distinctCount {
			result += (n - right + 1)
		}
	}

	return result
}
