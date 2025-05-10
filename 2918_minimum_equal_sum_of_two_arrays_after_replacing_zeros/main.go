package main

func minSum(nums1 []int, nums2 []int) int64 {
	nums1Sum := int64(0)
	nums2Sum := int64(0)
	nums1Zero := int64(0)
	nums2Zero := int64(0)

	for _, num := range nums1 {
		if num == 0 {
			nums1Zero++
			nums1Sum += 1
		} else {
			nums1Sum += int64(num)
		}
	}

	for _, num := range nums2 {
		if num == 0 {
			nums2Zero++
			nums2Sum += 1
		} else {
			nums2Sum += int64(num)
		}
	}

	if nums1Sum > nums2Sum && nums2Zero == 0 {
		return -1
	}

	if nums2Sum > nums1Sum && nums1Zero == 0 {
		return -1
	}

	return max(nums1Sum, nums2Sum)
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
