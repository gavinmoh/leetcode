package main

import "math"

// use this as key for hashmap
type Key struct {
	i         int
	k         int
	prev      string
	prevCount int
}

func getLengthOfOptimalCompression(s string, k int) int {
	cache := make(map[Key]int)

	var count func(i int, k int, prev string, prevCount int) int
	count = func(i int, k int, prev string, prevCount int) int {
		key := Key{i: i, k: k, prev: prev, prevCount: prevCount}
		if cachedResult, ok := cache[key]; ok {
			return cachedResult
		}
		if k < 0 {
			return math.MaxInt
		}
		if i == len(s) {
			return 0
		}

		var res int
		if string(s[i]) == prev {
			var incr int
			switch prevCount {
			case 1, 9, 99:
				incr = 1
			default:
				incr = 0
			}
			res = count(i+1, k, prev, prevCount+1) + incr
		} else {
			// get the minimum of not delete and delete current character
			res = min(count(i+1, k-1, prev, prevCount), count(i+1, k, string(s[i]), 1)+1)
		}

		cache[key] = res
		return res
	}

	return count(0, k, "", 0)

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
