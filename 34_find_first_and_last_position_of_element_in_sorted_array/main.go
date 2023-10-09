package main

func searchRange(nums []int, target int) []int {
	firstPosition := -1
	lastPosition := -1

	for i, v := range nums {
		if v == target {
			if firstPosition == -1 {
				firstPosition = i
			}
			lastPosition = i
		}
	}

	return []int{firstPosition, lastPosition}
}
