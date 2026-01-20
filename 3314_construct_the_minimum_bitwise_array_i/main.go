package main

func minBitwiseArray(nums []int) []int {
	result := make([]int, len(nums))
	for i := range result {
		result[i] = -1
	}

	for i, num := range nums {
		for j := num; j >= 0; j-- {
			if j|(j+1) == num {
				result[i] = j
			}
		}
	}

	return result
}
