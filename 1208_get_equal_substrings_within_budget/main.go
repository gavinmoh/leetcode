package main

func equalSubstring(s string, t string, maxCost int) int {
	n := len(s)
	costs := make([]int, n)
	for i := 0; i < n; i++ {
		costs[i] = abs(int(s[i]) - int(t[i]))
	}

	maxLength := 0
	currentCost := 0
	left := 0
	for right, cost := range costs {
		currentCost += cost
		length := right - left + 1
		if currentCost <= maxCost {
			if length > maxLength {
				maxLength = length
			}
		} else {
			currentCost -= costs[left]
			left++
		}
	}

	return maxLength
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
