package main

func maximumSubarraySum(nums []int, k int) int64 {
	maxSum := int64(0)
	runningSum := int64(0)
	numIndex := map[int]int{}
	left, right := 0, 0

	for right < len(nums) {
		num := nums[right]
		last := -1
		if _, ok := numIndex[num]; ok {
			last = numIndex[num]
		}

		for left <= last || right-left+1 > k {
			runningSum -= int64(nums[left])
			left++
		}

		numIndex[num] = right
		runningSum += int64(num)

		if right-left+1 == k && runningSum > maxSum {
			maxSum = runningSum
		}

		right++
	}

	return maxSum
}
