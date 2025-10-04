package main

func maxArea(height []int) int {
	maxArea := 0
	left, right := 0, len(height)-1

	for left < right {
		distance := right - left
		area := distance * min(height[left], height[right])
		maxArea = max(maxArea, area)

		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}

	return maxArea
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
