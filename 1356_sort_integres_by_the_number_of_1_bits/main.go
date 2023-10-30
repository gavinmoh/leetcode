package main

import (
	"sort"
)

func sortByBits(arr []int) []int {
	cache := make(map[int]int, len(arr))

	var countBits = func(num int) int {
		count := 0
		for num > 0 {
			if num&1 == 1 {
				count++
			}
			num >>= 1
		}
		return count
	}

	sort.Slice(arr, func(i, j int) bool {
		if _, ok := cache[arr[i]]; !ok {
			cache[arr[i]] = countBits(arr[i])
		}
		if _, ok := cache[arr[j]]; !ok {
			cache[arr[j]] = countBits(arr[j])
		}
		if cache[arr[i]] == cache[arr[j]] {
			return arr[i] < arr[j]
		}
		return cache[arr[i]] < cache[arr[j]]
	})

	return arr
}
