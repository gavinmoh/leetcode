package main

import "sort"

func maximumValueSum(nums []int, k int, _ [][]int) int64 {
	n := len(nums)
	delta := make([]int, n)
	sum := int64(0)

	// calculate the differences between num and (num XOR k)
	for i, num := range nums {
		sum += int64(num)
		delta[i] = (num ^ k) - num
	}
	sort.Ints(delta)

	// since we have to XOR two nums at the same time,
	// we take deltas and check if the XOR operation will
	// result in increment of the total sum
	// if yes, we add it to the sum, otherwise break
	for i := n - 1; i > 0; i -= 2 {
		deltaSum := delta[i] + delta[i-1]
		if deltaSum > 0 {
			sum += int64(deltaSum)
		} else {
			break
		}
	}

	return sum
}
