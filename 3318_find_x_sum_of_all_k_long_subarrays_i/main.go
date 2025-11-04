package main

import "sort"

func findXSum(nums []int, k int, x int) []int {
	n := len(nums)
	result := []int{}

	for i := 0; i < n-k+1; i++ {
		subarr := nums[i : i+k]

		freq := map[int]int{}
		for _, num := range subarr {
			freq[num]++
		}

		values := []int{}
		for num, _ := range freq {
			values = append(values, num)
		}

		sort.SliceStable(values, func(i, j int) bool {
			x, y := values[i], values[j]
			if freq[x] == freq[y] {
				return x > y
			}

			return freq[x] > freq[y]
		})

		xSum := 0
		for j := 0; j < min(x, len(values)); j++ {
			xSum += values[j] * freq[values[j]]
		}

		result = append(result, xSum)
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
