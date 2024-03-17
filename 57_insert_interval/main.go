package main

func insert(intervals [][]int, newInterval []int) [][]int {
	lefts := [][]int{}
	rights := [][]int{}

	for _, interval := range intervals {
		if interval[1] < newInterval[0] {
			lefts = append(lefts, interval)
		} else if interval[0] > newInterval[1] {
			rights = append(rights, interval)
		} else {
			newInterval = []int{
				min(newInterval[0], interval[0]),
				max(newInterval[1], interval[1]),
			}
		}
	}

	result := append(lefts, newInterval)
	result = append(result, rights...)
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
