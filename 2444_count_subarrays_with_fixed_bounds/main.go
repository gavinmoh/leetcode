package main

func countSubarrays(nums []int, minK int, maxK int) int64 {
	prevMin := -1
	prevMax := -1
	boundary := -1
	result := int64(0)

	for i, num := range nums {
		if num < minK || num > maxK {
			boundary = i
			continue
		}

		if num == minK {
			prevMin = i
		}
		if num == maxK {
			prevMax = i
		}

		result += int64(max(0, min(prevMin, prevMax)-boundary))
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
