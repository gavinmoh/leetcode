package main

func resultsArray(nums []int, k int) []int {
	result := []int{}
	for i := 0; i <= len(nums)-k; i++ {
		subarray := nums[i : i+k]
		if isConsecutive(subarray) {
			result = append(result, subarray[k-1])
		} else {
			result = append(result, -1)
		}
	}
	return result
}

func isConsecutive(nums []int) bool {
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] != 1 {
			return false
		}
	}
	return true
}
