package main

import "sort"

func findLHS(nums []int) int {
	freq := map[int]int{}
	for _, num := range nums {
		freq[num]++
	}

	uniq := []int{}
	for num := range freq {
		uniq = append(uniq, num)
	}
	sort.Ints(uniq)

	maxLength := 0
	for i := 1; i < len(uniq); i++ {
		if abs(uniq[i]-uniq[i-1]) == 1 {
			length := freq[uniq[i]] + freq[uniq[i-1]]
			if length > maxLength {
				maxLength = length
			}
		}
	}

	return maxLength
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
