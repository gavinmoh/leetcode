package main

import "sort"

func fullBloomFlowers(flowers [][]int, people []int) []int {
	flowersStart := make([]int, len(flowers))
	flowersEnd := make([]int, len(flowers))
	for i, flower := range flowers {
		flowersStart[i] = flower[0]
		flowersEnd[i] = flower[1]
	}

	sort.Ints(flowersStart)
	sort.Ints(flowersEnd)

	result := make([]int, len(people))
	for i, day := range people {
		start := sort.SearchInts(flowersEnd, day)
		end := sort.SearchInts(flowersStart, day+1)
		result[i] = end - start
	}

	return result

}
