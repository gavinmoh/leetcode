package main

func maximumTripletValue(nums []int) int64 {
	n := len(nums)
	maxValue := int64(0)

	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				value := int64(nums[i]-nums[j]) * int64(nums[k])
				if value > maxValue {
					maxValue = value
				}
			}
		}
	}

	return maxValue
}
