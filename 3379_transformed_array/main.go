package main

func constructTransformedArray(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	for i, num := range nums {
		j := (i + num) % n
		if j < 0 {
			j = n - (abs(j) % n)
		}
		result[i] = nums[j]
	}

	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
