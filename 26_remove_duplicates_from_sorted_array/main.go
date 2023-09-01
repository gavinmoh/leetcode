package main

func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	k := 1 // k is the index of the last unique element
	for i := 1; i < len(nums); i++ {
		// when we find a new unique element, we put it in the kth position
		// and increment k; we ignore all the duplicates
		if nums[i] != nums[i-1] {
			nums[k] = nums[i]
			k++
		}
	}

	return k
}
