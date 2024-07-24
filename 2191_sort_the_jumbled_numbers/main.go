package main

import "sort"

func sortJumbled(mapping []int, nums []int) []int {
	mapped := make(map[int]int)
	for _, num := range nums {
		mappedNum, n := []int{}, num
		for n > 0 {
			digit := n % 10
			mappedNum = append(mappedNum, mapping[digit])
			n = n / 10
		}
		if len(mappedNum) == 0 {
			mappedNum = append(mappedNum, mapping[num])
		}
		output := mappedNum[0]
		for i, acc := 1, 10; i < len(mappedNum); i++ {
			output += mappedNum[i] * acc
			acc = acc * 10
		}
		mapped[num] = output
	}

	sort.SliceStable(nums, func(i, j int) bool {
		return mapped[nums[i]] < mapped[nums[j]]
	})

	return nums
}
