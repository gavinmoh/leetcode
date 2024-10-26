package main

import "sort"

func longestSquareStreak(nums []int) int {
	sort.Ints(nums)
	numsMap := map[int]bool{}

	for _, num := range nums {
		numsMap[num] = true
	}

	maxStreak := -1
	for _, num := range nums {
		currentStreak := 1
		for {
			num *= num
			if _, ok := numsMap[num]; !ok {
				break
			}
			currentStreak++
		}
		if currentStreak > 1 && currentStreak > maxStreak {
			maxStreak = currentStreak
		}
	}

	return maxStreak
}
