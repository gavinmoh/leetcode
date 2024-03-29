package main

func countSubarrays(nums []int, k int) int64 {
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}

	count := int64(0)
	indices := []int{}
	for i, num := range nums {
		if num == max {
			indices = append(indices, i)
		}

		freq := len(indices)
		if freq >= k {
			count += int64(indices[freq-k] + 1)
		}
	}
	return count
}
