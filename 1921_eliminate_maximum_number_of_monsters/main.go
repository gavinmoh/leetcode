package main

import (
	"sort"
)

func eliminateMaximum(dist []int, speed []int) int {
	// the key to this problem is to calculate the time in terms of float and not int
	eta := make([]float32, len(dist))
	for i := 0; i < len(dist); i++ {
		eta[i] = float32(dist[i]) / float32(speed[i])
	}
	sort.Slice(eta, func(i, j int) bool {
		return eta[i] < eta[j]
	})

	count := 0
	for i := 0; i < len(eta); i++ {
		if eta[i] <= float32(i) {
			break
		}
		count++
	}
	return count
}
