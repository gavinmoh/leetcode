package main

func intersection(nums1 []int, nums2 []int) []int {
	num1Map := make(map[int]bool)
	for _, num := range nums1 {
		num1Map[num] = true
	}

	resultMap := make(map[int]bool)
	for _, num := range nums2 {
		if num1Map[num] {
			resultMap[num] = true
		}
	}

	result := []int{}
	for num := range resultMap {
		result = append(result, num)
	}

	return result
}
