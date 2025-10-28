package main

func countValidSelections(nums []int) int {
	n := len(nums)
	leftSum, rightSum := make([]int, n), make([]int, n)

	leftSum[0] = nums[0]
	for i := 1; i < n; i++ {
		leftSum[i] += nums[i] + leftSum[i-1]
	}

	rightSum[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		rightSum[i] += nums[i] + rightSum[i+1]
	}

	count := 0
	for i := 0; i < n; i++ {
		if nums[i] != 0 {
			continue
		}

		if leftSum[i] == rightSum[i] {
			count += 2
		}

		if leftSum[i] == rightSum[i]-1 {
			count += 1
		}

		if leftSum[i]-1 == rightSum[i] {
			count += 1
		}
	}

	return count
}
