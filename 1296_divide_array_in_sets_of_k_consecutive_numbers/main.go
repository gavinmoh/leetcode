package main

import "math"

func isPossibleDivide(nums []int, k int) bool {
	if len(nums)%k != 0 {
		return false
	}

	numsMap := make(map[int]int)
	min, max := math.MaxInt, math.MinInt
	for _, num := range nums {
		numsMap[num]++
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}

	start := min
	for group := 0; group < len(nums)/k; group++ {
		nextStart, nextStartFound := 0, false
		for i := 0; i < k; i++ {
			current := start + i
			if count, ok := numsMap[current]; ok && count > 0 {
				numsMap[current]--
				if count > 1 && !nextStartFound {
					nextStart = current
					nextStartFound = true
				}
			} else {
				return false
			}
		}

		if nextStartFound {
			start = nextStart
		} else {
			// find the next start manually
			for i := start + k; i <= max-k+1; i++ {
				if count, ok := numsMap[i]; ok && count > 0 {
					start = i
					break
				}
			}
		}
	}

	return true
}
