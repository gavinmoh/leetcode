package main

func majorityElement(nums []int) int {
	n := len(nums)
	majority := n / 2
	freq := make(map[int]int)

	for _, num := range nums {
		freq[num]++
		if freq[num] > majority {
			return num
		}
	}

	return 0
}
