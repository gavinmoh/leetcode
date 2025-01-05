package main

func xorAllNums(nums1 []int, nums2 []int) int {
	m, n := len(nums1), len(nums2)
	result := 0
	if n%2 != 0 {
		for _, num := range nums1 {
			result ^= num
		}
	}

	if m%2 != 0 {
		for _, num := range nums2 {
			result ^= num
		}
	}

	return result
}
