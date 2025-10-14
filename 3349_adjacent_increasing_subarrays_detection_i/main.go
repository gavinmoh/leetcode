package main

func hasIncreasingSubarrays(nums []int, k int) bool {
	if k == 1 {
		return true
	}

	n := len(nums)
	windowSize := k * 2
	for i := 0; i <= n-windowSize; i++ {
		subarray1 := nums[i : i+k]
		subarray2 := nums[i+k : i+windowSize]
		if isStrictlyIncreasing(subarray1) && isStrictlyIncreasing(subarray2) {
			return true
		}
	}

	return false
}

func isStrictlyIncreasing(arr []int) bool {
	n := len(arr)
	for i := 1; i < n; i++ {
		if arr[i] <= arr[i-1] {
			return false
		}
	}

	return true
}
