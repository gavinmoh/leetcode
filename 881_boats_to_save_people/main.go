package main

import "sort"

func numRescueBoats(people []int, limit int) int {
	boats := 0
	sort.Ints(people)

	for left, right := 0, len(people)-1; left <= right; right-- {
		limitLeft := limit - people[right]
		if people[left] <= limitLeft {
			left++
		}
		boats++
	}

	return boats
}
