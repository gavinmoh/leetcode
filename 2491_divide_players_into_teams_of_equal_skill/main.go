package main

import "sort"

func dividePlayers(skill []int) int64 {
	sort.Ints(skill)

	n := len(skill)
	sum := skill[0] + skill[n-1]
	chemistry := int64(skill[0] * skill[n-1])

	left, right := 1, n-2
	for left < right {
		currentSum := skill[left] + skill[right]
		if currentSum != sum {
			return -1
		}
		chemistry += int64(skill[left] * skill[right])
		left++
		right--
	}

	return chemistry
}
