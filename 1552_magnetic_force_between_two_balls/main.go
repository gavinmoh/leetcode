package main

import "sort"

func maxDistance(position []int, m int) int {
	n := len(position)
	canPlace := func(dist int) bool {
		count := 1
		for i := 0; i < n; {
			j := i
			for j < n && position[j]-position[i] < dist {
				j++
			}
			if j == n {
				break
			} else {
				count++
				i = j
			}
			if count == m {
				return true
			}
		}
		return false
	}

	sort.Ints(position)
	left, right := 1, position[n-1]-position[0]
	for left < right {
		mid := right - (right-left)/2
		if canPlace(mid) {
			left = mid
		} else {
			right = mid - 1
		}
	}

	return left
}
