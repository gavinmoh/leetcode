package main

import (
	"sort"
)

type Item struct {
	Index int
	Max   int
	Min   int
}

func maxDistance(arrays [][]int) int {
	sortedByMax := []*Item{}
	sortedByMin := []*Item{}
	for i, array := range arrays {
		item := &Item{
			Index: i,
			Max:   array[len(array)-1],
			Min:   array[0],
		}
		sortedByMax = append(sortedByMax, item)
		sortedByMin = append(sortedByMin, item)
	}

	sort.SliceStable(sortedByMax, func(i, j int) bool {
		return sortedByMax[i].Max > sortedByMax[j].Max
	})
	sort.SliceStable(sortedByMin, func(i, j int) bool {
		return sortedByMin[i].Min < sortedByMin[j].Min
	})

	if sortedByMax[0].Index == sortedByMin[0].Index {
		left := sortedByMax[0].Max - sortedByMin[1].Min
		right := sortedByMax[1].Max - sortedByMin[0].Min
		return max(left, right)
	}

	return sortedByMax[0].Max - sortedByMin[0].Min
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
