package main

func minimumOperations(nums []int) int {
	count := 0

	for _, num := range nums {
		remainder := num % 3
		if remainder == 0 {
			continue
		}

		count++
	}

	return count
}
