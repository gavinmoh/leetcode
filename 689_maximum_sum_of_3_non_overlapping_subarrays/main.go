package main

func maxSumOfThreeSubarrays(nums []int, k int) []int {
	n := len(nums)

	prefixSum := make([]int, n+1)
	for i, num := range nums {
		prefixSum[i+1] = num + prefixSum[i]
	}

	bestSum := make([][]int, 4)
	for i := 0; i < 4; i++ {
		bestSum[i] = make([]int, n+1)
	}

	bestIndex := make([][]int, 4)
	for i := 0; i < 4; i++ {
		bestIndex[i] = make([]int, n+1)
	}

	for t := 1; t < 4; t++ {
		for i := k * t; i < n+1; i++ {
			currentSum := prefixSum[i] - prefixSum[i-k] + bestSum[t-1][i-k]

			if currentSum > bestSum[t][i-1] {
				bestSum[t][i] = currentSum
				bestIndex[t][i] = i - k
			} else {
				bestSum[t][i] = bestSum[t][i-1]
				bestIndex[t][i] = bestIndex[t][i-1]
			}
		}
	}

	result := make([]int, 3)
	end := n
	for t := 3; t > 0; t-- {
		result[t-1] = bestIndex[t][end]
		end = result[t-1]
	}

	return result
}
