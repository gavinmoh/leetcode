package main

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	indices := make(map[int]int, len(nums1)) // nums1 value -> position
	for i, num := range nums1 {
		indices[num] = i
	}
	// initialize result to -1
	result := make([]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		result[i] = -1
	}

	stack := []int{}
	for _, num := range nums2 {
		for len(stack) > 0 && num > stack[len(stack)-1] {
			// pop the value
			val := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			index := indices[val]
			result[index] = num
		}
		if _, ok := indices[num]; ok {
			stack = append(stack, num)
		}
	}

	return result
}
