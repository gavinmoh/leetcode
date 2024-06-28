package main

import "sort"

func maximumImportance(n int, roads [][]int) int64 {
	degrees := make([]int, n)
	for _, road := range roads {
		degrees[road[0]]++
		degrees[road[1]]++
	}
	sort.Ints(degrees)

	value := int64(1)
	total := int64(0)
	for _, degree := range degrees {
		total += value * int64(degree)
		value++
	}

	return total
}
