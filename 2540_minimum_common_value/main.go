package main

func getCommon(nums1 []int, nums2 []int) int {
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		num1, num2 := nums1[i], nums2[j]
		if num1 == num2 {
			return num1
		} else if num1 > num2 {
			j++
		} else {
			i++
		}
	}

	return -1
}
