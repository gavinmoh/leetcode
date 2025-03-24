package main

import "sort"

func countDays(days int, meetings [][]int) int {
	sort.SliceStable(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})

	merged := [][]int{}
	currentStart, currentEnd := meetings[0][0], meetings[0][1]
	for i := 1; i < len(meetings); i++ {
		nextStart, nextEnd := meetings[i][0], meetings[i][1]
		if currentEnd > nextEnd {
			continue
		}

		if currentEnd >= nextStart {
			currentEnd = nextEnd
			continue
		}

		merged = append(merged, []int{currentStart, currentEnd})
		currentStart, currentEnd = nextStart, nextEnd
	}
	merged = append(merged, []int{currentStart, currentEnd})

	currentDay, idx := 0, 0
	count := 0
	for currentDay <= days && idx < len(merged) {
		start, end := merged[idx][0], merged[idx][1]
		if currentDay < start {
			count += start - currentDay - 1
		}
		currentDay = end
		idx++
	}
	if currentDay < days {
		count += days - currentDay
	}

	return count
}
