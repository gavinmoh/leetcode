package main

func canSortArray(nums []int) bool {
	setBits := map[int]int{}
	for _, num := range nums {
		setBits[num] = countSetBits(num)
	}

	// bubble sort
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] <= nums[j+1] {
				continue
			}
			if setBits[nums[j]] != setBits[nums[j+1]] {
				return false
			}
			nums[j], nums[j+1] = nums[j+1], nums[j]
		}
	}

	return true
}

func countSetBits(num int) int {
	count := 0
	for num > 0 {
		if num&1 == 1 {
			count++
		}
		num >>= 1
	}

	return count
}
