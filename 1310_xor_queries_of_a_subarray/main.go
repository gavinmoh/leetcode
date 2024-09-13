package main

func xorQueries(arr []int, queries [][]int) []int {
	prefixSum := make([]int, len(arr))
	prefixSum[0] = arr[0]
	for i := 1; i < len(arr); i++ {
		prefixSum[i] = prefixSum[i-1] ^ arr[i]
	}

	answer := make([]int, len(queries))
	for i, query := range queries {
		left, right := query[0], query[1]
		if left == 0 {
			answer[i] = prefixSum[right]
		} else if (right - left) == 0 {
			answer[i] = arr[right]
		} else {
			answer[i] = prefixSum[right] ^ prefixSum[left-1]
		}
	}

	return answer
}
