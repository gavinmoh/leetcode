package main

import "sort"

func putMarbles(weights []int, k int) int64 {
	n := len(weights)
	pairWeights := []int{}
	for i := 0; i < n-1; i++ {
		pairWeights = append(pairWeights, weights[i]+weights[i+1])
	}

	sort.Ints(pairWeights)

	answer := int64(0)
	for i := 0; i < k-1; i++ {
		answer += int64(pairWeights[n-2-i] - pairWeights[i])
	}

	return answer
}
