package main

func longestNiceSubarray(nums []int) int {
	left, maxLength, usedBits := 0, 0, 0
	for right := 0; right < len(nums); right++ {
		for usedBits&nums[right] != 0 {
			usedBits ^= nums[left]
			left++
		}

		usedBits |= nums[right]
		length := right - left + 1
		if length > maxLength {
			maxLength = length
		}
	}

	return maxLength
}
