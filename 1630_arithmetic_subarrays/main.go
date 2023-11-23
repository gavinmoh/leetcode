package main

import "sort"

func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	n := len(l)
	output := make([]bool, n)

outer:
	for i := 0; i < n; i++ {
		originalArray := nums[l[i] : r[i]+1]
		subarray := make([]int, len(originalArray))
		copy(subarray, originalArray)

		sort.Ints(subarray)
		diff := subarray[1] - subarray[0]

		for j := 1; j < len(subarray)-1; j++ {
			if diff != (subarray[j+1] - subarray[j]) {
				continue outer
			}
		}
		output[i] = true
	}

	return output

}
