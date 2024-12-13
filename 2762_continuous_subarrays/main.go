package main

import "math"

func continuousSubarrays(nums []int) int64 {
	freq := map[int]int{}
	left := 0
	count := int64(0)

	for right := 0; right < len(nums); right++ {
		freq[nums[right]]++

		for max(freq)-min(freq) > 2 {
			freq[nums[left]]--
			if freq[nums[left]] == 0 {
				delete(freq, nums[left])
			}
			left++
		}

		count += int64(right - left + 1)
	}

	return count
}

func max(freq map[int]int) int {
	maxNum := math.MinInt
	for num, _ := range freq {
		if num > maxNum {
			maxNum = num
		}
	}

	return maxNum
}

func min(freq map[int]int) int {
	minNum := math.MaxInt
	for num, _ := range freq {
		if num < minNum {
			minNum = num
		}
	}

	return minNum
}
