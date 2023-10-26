package main

import (
	"sort"
)

func numFactoredBinaryTrees(arr []int) int {
	// dp[i] is the number of binary trees that can be formed with arr[i] as root
	// since a single value can be a binary tree on its own, we initialize dp[i] to 1
	dp := make(map[int]int, len(arr))
	for _, i := range arr {
		dp[i] = 1
	}

	// sort the array
	sort.Ints(arr)

	// for each value in the array, we check if it can be formed by multiplying two other values in the array
	// if it can, we add the number of binary trees that can be formed with those two values as roots
	// to the number of binary trees that can be formed with the current value as root
	// we do this for all values in the array
	for i := 0; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			if arr[i]%arr[j] == 0 {
				k := arr[i] / arr[j]
				if _, ok := dp[k]; ok {
					dp[arr[i]] += dp[arr[j]] * dp[k]
				}
			}
		}
	}

	var ans int
	for _, v := range dp {
		ans += v
	}

	return ans % (1e9 + 7)
}
