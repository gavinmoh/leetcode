package main

func missingNumber(nums []int) int {
	n := len(nums) + 1
	numsMap := make(map[int]bool, n-1)
	for _, num := range nums {
		numsMap[num] = true
	}
	for i := 0; i < n; i++ {
		if !numsMap[i] {
			return i
		}
	}

	return -1
}
