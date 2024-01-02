package main

func findMatrix(nums []int) [][]int {
	processed := make(map[int]int)
	result := [][]int{}

	for _, num := range nums {
		if _, ok := processed[num]; !ok {
			processed[num] = 1
		} else {
			processed[num] += 1
		}

		if len(result) < processed[num] {
			result = append(result, []int{num})
		} else {
			result[processed[num]-1] = append(result[processed[num]-1], num)
		}
	}

	return result
}
