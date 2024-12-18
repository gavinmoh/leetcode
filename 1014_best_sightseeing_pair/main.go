package main

func maxScoreSightseeingPair(values []int) int {
	maxValue := values[0]
	result := -1
	for i := 1; i < len(values); i++ {
		current := values[i] - i
		result = max(result, maxValue+current)
		maxValue = max(maxValue, values[i]+i)
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
