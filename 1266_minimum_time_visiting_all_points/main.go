package main

func minTimeToVisitAllPoints(points [][]int) int {
	output := 0
	currentPoint := points[0]

	for i := 1; i < len(points); i++ {
		nextPoint := points[i]
		// distance of x = x2 - x1
		// distnace of y = y2 - y1
		dx := abs(nextPoint[0] - currentPoint[0])
		dy := abs(nextPoint[1] - currentPoint[1])
		diff := abs(dx - dy)

		// if they're the same, then we can move diagonally all the way
		if diff == 0 {
			output += dx
		} else {
			output += max(dx, dy)
		}

		currentPoint = nextPoint
	}

	return output
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
