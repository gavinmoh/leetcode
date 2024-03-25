package main

func findDuplicates(nums []int) []int {
	appearance := make([]bool, 100001)
	result := []int{}
	for _, num := range nums {
		if appearance[num] {
			result = append(result, num)
		} else {
			appearance[num] = true
		}
	}

	return result
}
