package main

func minCost(colors string, neededTime []int) int {
	time := 0
	runningSum := neededTime[0]
	maxValue := neededTime[0]

	for i := 1; i < len(colors); i++ {
		if colors[i] == colors[i-1] {
			runningSum += neededTime[i]
			maxValue = max(maxValue, neededTime[i])
		} else {
			time += (runningSum - maxValue)
			runningSum = neededTime[i]
			maxValue = neededTime[i]
		}
	}
	time += (runningSum - maxValue)

	return time
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
