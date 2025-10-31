package main

func getSneakyNumbers(nums []int) []int {
	n := len(nums)
	freq := make([]int, n-2)
	for _, num := range nums {
		freq[num]++
	}

	result := []int{}
	for num, count := range freq {
		if count > 1 {
			result = append(result, num)
		}
	}

	return result
}
