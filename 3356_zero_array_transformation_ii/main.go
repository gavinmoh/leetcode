package main

func minZeroArray(nums []int, queries [][]int) int {
	n := len(nums)
	totalSum := 0
	k := 0
	differenceArray := make([]int, n+1)

	for i := 0; i < n; i++ {
		for totalSum+differenceArray[i] < nums[i] {
			k++

			if k > len(queries) {
				return -1
			}

			query := queries[k-1]
			left, right, val := query[0], query[1], query[2]

			if right >= i {
				differenceArray[max(left, i)] += val
				differenceArray[right+1] -= val
			}
		}
		totalSum += differenceArray[i]
	}

	return k
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
