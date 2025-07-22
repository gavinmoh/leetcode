package main

func maximumUniqueSubarray(nums []int) int {
	maxScore, currentScore := 0, 0
	count := map[int]int{}

	for left, right := 0, 0; right < len(nums); right++ {
		count[nums[right]]++
		currentScore += nums[right]

		for count[nums[right]] > 1 {
			count[nums[left]]--
			currentScore -= nums[left]
			left++
		}

		maxScore = max(maxScore, currentScore)
	}

	return maxScore
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
