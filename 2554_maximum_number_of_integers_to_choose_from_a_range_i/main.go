package main

import "sort"

func maxCount(banned []int, n int, maxSum int) int {
	sort.Ints(banned)

	isBanned := func(num int) bool {
		left, right := 0, len(banned)-1
		for left <= right {
			mid := (left + right) / 2
			if banned[mid] == num {
				return true
			} else if banned[mid] > num {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}

		return false
	}

	count := 0
	for i := 1; i <= n; i++ {
		if isBanned(i) {
			continue
		}
		maxSum -= i
		if maxSum < 0 {
			break
		}
		count++
	}

	return count
}
