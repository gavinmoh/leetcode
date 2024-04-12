package main

func trap(height []int) int {
	n := len(height)
	leftMax := make([]int, n)
	rightMax := make([]int, n)

	for i, currentMax := 0, 0; i < n; i++ {
		if height[i] > currentMax {
			currentMax = height[i]
		}
		leftMax[i] = currentMax
	}

	for i, currentMax := n-1, 0; i >= 0; i-- {
		if height[i] > currentMax {
			currentMax = height[i]
		}
		rightMax[i] = currentMax
	}

	sum := 0
	for i, h := range height {
		boundary := min(leftMax[i], rightMax[i])
		sum += boundary - h
	}

	return sum
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
