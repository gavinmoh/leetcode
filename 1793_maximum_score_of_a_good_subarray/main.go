package main

func maximumScore(nums []int, k int) int {
	left := k
	right := k
	currentMin := nums[k]
	maxScore := currentMin * (right - left + 1)

	for left >= 0 && right < len(nums) {
		currentMin = min(currentMin, min(nums[left], nums[right]))
		score := currentMin * (right - left + 1)
		if score > maxScore {
			maxScore = score
		}

		if left == 0 && right == len(nums)-1 {
			break
		}

		if left == 0 {
			right++
			continue
		}

		if right == len(nums)-1 {
			left--
			continue
		}

		if nums[left-1] > nums[right+1] {
			left--
		} else {
			right++
		}
	}

	return maxScore
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
