package main

func longestBalanced(nums []int) int {
	maxLength := 0

	for left := 0; left < len(nums); left++ {
		freq := map[int]bool{}
		odd, even := 0, 0
		for right := left; right < len(nums); right++ {
			if _, ok := freq[nums[right]]; !ok {
				freq[nums[right]] = true
				if nums[right]%2 == 0 {
					even++
				} else {
					odd++
				}
			}

			if odd != even {
				continue
			}

			length := right - left + 1
			if length > maxLength {
				maxLength = length
			}
		}
	}

	return maxLength
}
