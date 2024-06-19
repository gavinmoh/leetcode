package main

import "math"

func minDays(bloomDay []int, m int, k int) int {
	required := m * k
	if required > len(bloomDay) {
		return -1
	}

	canMakeBouquets := func(day int) bool {
		count := 0
		consecutive := 0
		for i := 0; i < len(bloomDay); i++ {
			if bloomDay[i] > day {
				consecutive = 0
			} else {
				consecutive += 1
				if consecutive == k {
					count += 1
					consecutive = 0

					if count >= m {
						break
					}
				}
			}
		}

		return count == m
	}

	left := 1
	right := math.MaxInt
	for left < right {
		mid := left + ((right - left) / 2)
		if canMakeBouquets(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}
