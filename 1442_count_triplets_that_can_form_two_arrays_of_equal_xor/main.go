package main

func countTriplets(arr []int) int {
	n := len(arr)
	prefixSum := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefixSum[i+1] = prefixSum[i] ^ arr[i]
	}

	count := 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			for k := j; k < n; k++ {
				a := prefixSum[j] ^ prefixSum[i]
				b := prefixSum[k+1] ^ prefixSum[j]
				if a == b {
					count++
				}
			}
		}
	}

	return count
}
