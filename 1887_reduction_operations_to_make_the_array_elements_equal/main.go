package main

import "sort"

func reductionOperations(nums []int) int {
	n := len(nums)
	sort.Ints(nums)
	smallest := nums[0]
	largest := nums[n-1]
	uniqueNums := []int{}
	freqMap := make(map[int]int)

	for i := n - 1; i >= 0; i-- {
		num := nums[i]
		if freqMap[num] == 0 {
			uniqueNums = append(uniqueNums, num)
		}
		freqMap[num] += 1
	}

	if smallest == largest {
		return 0
	}

	uniqCount := len(uniqueNums) - 1
	output := 0
	for i := 0; ; i++ {
		if uniqCount <= 0 {
			break
		}
		num := uniqueNums[i]
		count := freqMap[num]
		output += (count * uniqCount)
		uniqCount--
	}
	return output
}
