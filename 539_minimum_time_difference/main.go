package main

import (
	"math"
	"sort"
	"strconv"
)

func findMinDifference(timePoints []string) int {
	minutes := []int{}
	for _, point := range timePoints {
		hour, _ := strconv.Atoi(point[:2])
		minute, _ := strconv.Atoi(point[3:])
		minutes = append(minutes, hour*60+minute)
	}
	sort.Ints(minutes)
	minutes = append(minutes, minutes[0]+1440)

	minDiff := math.MaxInt
	for i := 0; i < len(timePoints); i++ {
		diff := minutes[i+1] - minutes[i]
		if diff < minDiff {
			minDiff = diff
		}
	}

	return minDiff
}
