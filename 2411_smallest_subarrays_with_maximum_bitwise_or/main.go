package main

func smallestSubarrays(nums []int) []int {
	n := len(nums)
	answer := make([]int, n)
	for i := range nums {
		answer[i] = 1
	}

	for i, num := range nums {
		for j := i - 1; j >= 0 && (nums[j]|num) != nums[j]; j-- {
			answer[j] = i - j + 1
			nums[j] |= num
		}
	}

	return answer
}
