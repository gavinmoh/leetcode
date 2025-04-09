package main

func minOperations(nums []int, k int) int {
	uniqMap := map[int]bool{}
	for _, num := range nums {
		uniqMap[num] = true
	}

	operation := 0
	for num := range uniqMap {
		if num < k {
			return -1
		} else if num > k {
			operation++
		}
	}

	return operation
}
