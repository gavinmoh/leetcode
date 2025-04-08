package main

func minimumOperations(nums []int) int {
	freq := map[int]int{}
	duplicated := map[int]bool{}

	for _, num := range nums {
		freq[num]++
		if freq[num] == 2 {
			duplicated[num] = true
		}
	}

	operation := 0
	for len(duplicated) > 0 {
		operation++

		if len(nums) < 3 {
			break
		}

		for i := 0; i < 3; i++ {
			num := nums[i]
			freq[num]--
			if freq[num] == 1 {
				delete(duplicated, num)
			}
		}
		nums = nums[3:]
	}

	return operation
}
