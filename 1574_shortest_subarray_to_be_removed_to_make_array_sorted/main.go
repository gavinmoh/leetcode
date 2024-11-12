package main

func findLengthOfShortestSubarray(arr []int) int {
	right := len(arr) - 1
	for right > 0 && arr[right] >= arr[right-1] {
		right--
	}

	result := right
	left := 0
	for left < right && (left == 0 || arr[left-1] <= arr[left]) {
		for right < len(arr) && arr[left] > arr[right] {
			right++
		}
		result = min(result, right-left-1)
		left++
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
