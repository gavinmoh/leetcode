package main

func CombinationSum4(nums []int, target int) int {
	result := make([]int, target+1)
	result[0] = 1

	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if i-num >= 0 {
				result[i] += result[i-num]
			}
		}
	}

	return result[target]
}
