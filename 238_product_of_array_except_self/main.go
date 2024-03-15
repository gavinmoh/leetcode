package main

func productExceptSelf(nums []int) []int {
	n := len(nums)
	lefts := make([]int, n)
	rights := make([]int, n)

	left := 1
	for i := 0; i < n; i++ {
		left *= nums[i]
		lefts[i] = left
	}

	right := 1
	for j := n - 1; j >= 0; j-- {
		right *= nums[j]
		rights[j] = right
	}

	result := make([]int, n)
	for k := 0; k < n; k++ {
		if k == 0 {
			result[k] = rights[k+1]
		} else if k+1 == n {
			result[k] = lefts[k-1]
		} else {
			result[k] = lefts[k-1] * rights[k+1]
		}
	}

	return result
}
