package main

func longestSubarray(nums []int) int {
	zeroCount, maxLength := 0, 0
	for left, right := 0, 0; right < len(nums); right++ {
		if nums[right] == 0 {
			zeroCount++
		}

		for zeroCount > 1 && left <= right {
			if nums[left] == 0 {
				zeroCount--
			}
			left++
		}

		length := right - left // must delete one element
		if length > maxLength {
			maxLength = length
		}
	}

	return maxLength
}
