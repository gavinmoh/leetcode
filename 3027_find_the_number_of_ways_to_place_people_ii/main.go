package main

import (
	"math"
	"sort"
)

func numberOfPairs(points [][]int) int {
	count := 0

	// sort points prioritizing upper left first
	sort.SliceStable(points, func(i, j int) bool {
		if points[i][0] == points[j][0] {
			return points[i][1] > points[j][1]
		}

		return points[i][0] < points[j][0]
	})

	for i, pointA := range points {
		x1, y1 := pointA[0], pointA[1]
		xMin := x1 - 1
		xMax := math.MaxInt32
		yMin := math.MinInt32
		yMax := y1 + 1

		for j := i + 1; j < len(points); j++ {
			pointB := points[j]
			x2, y2 := pointB[0], pointB[1]
			if x2 > xMin && x2 < xMax && y2 > yMin && y2 < yMax {
				count++
				xMin = x2
				yMin = y2
			}
		}
	}

	return count
}
