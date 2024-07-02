package main

func intersect(nums1 []int, nums2 []int) []int {
	nums2Map := make(map[int]int)
	for _, num := range nums2 {
		nums2Map[num]++
	}

	result := []int{}
	for _, num := range nums1 {
		if count, ok := nums2Map[num]; ok && count > 0 {
			result = append(result, num)
			nums2Map[num]--
		}
	}

	return result
}
