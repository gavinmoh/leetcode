package main

func GetMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func GetMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MaxArea(height []int) int {
	maxArea := 0

	// looping from the first
	for i := 0; i < len(height); {
		// looping from the last
		for j := len(height); j > i; {
			distance := j - i - 1
			min := GetMin(height[i], height[j-1])
			area := distance * min
			maxArea = GetMax(maxArea, area)
			if height[i] > height[j-1] { // if left is higher than right, move right
				j--
			} else {
				i++
			}
		}
	}
	return maxArea
}
