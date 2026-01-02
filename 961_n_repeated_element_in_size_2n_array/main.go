package main

func repeatedNTimes(nums []int) int {
	n := len(nums) / 2
	freq := map[int]int{}
	for _, num := range nums {
		freq[num]++
	}

	for num, count := range freq {
		if count == n {
			return num
		}
	}

	return -1
}
