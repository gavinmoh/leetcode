package main

import "sort"

func checkValidCuts(n int, rectangles [][]int) bool {
	x, y := [][]int{}, [][]int{}
	for _, rect := range rectangles {
		xStart, yStart, xEnd, yEnd := rect[0], rect[1], rect[2], rect[3]
		x = append(x, []int{xStart, xEnd})
		y = append(y, []int{yStart, yEnd})
	}

	sort.SliceStable(x, func(i, j int) bool {
		return x[i][0] < x[j][0]
	})

	sort.SliceStable(y, func(i, j int) bool {
		return y[i][0] < y[j][0]
	})

	// check for x directions
	count := 1
	currentEnd := x[0][1]
	for i := 1; i < len(x); i++ {
		nextStart, nextEnd := x[i][0], x[i][1]
		if currentEnd > nextStart {
			currentEnd = max(currentEnd, nextEnd)
			continue
		}
		count++
		currentEnd = nextEnd
	}

	if count >= 3 {
		return true
	}

	// check for y directions
	count = 1
	currentEnd = y[0][1]
	for i := 1; i < len(y); i++ {
		nextStart, nextEnd := y[i][0], y[i][1]
		if currentEnd > nextStart {
			currentEnd = max(currentEnd, nextEnd)
			continue
		}
		count++
		currentEnd = nextEnd
	}

	return count >= 3
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
