package main

import "sort"

func findMinArrowShots(points [][]int) int {
	// sort by ending points
	sort.SliceStable(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	previousEnd := points[0][1]
	count := 1
	for i := 1; i < len(points); i++ {
		// check if prev ending point is overlapped
		// if not overlapped, then we will need an extra arrow
		if previousEnd < points[i][0] {
			count++
			previousEnd = points[i][1]
		}
	}

	return count
}
