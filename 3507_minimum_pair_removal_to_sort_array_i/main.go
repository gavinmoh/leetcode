package main

func minimumPairRemoval(nums []int) int {
	nonDecreasing := func(arr []int) bool {
		for i := 1; i < len(arr); i++ {
			if arr[i] < arr[i-1] {
				return false
			}
		}
		return true
	}

	count := 0
	for !nonDecreasing(nums) {
		minIdx := 0
		minSum := nums[0] + nums[1]
		for i := 2; i < len(nums); i++ {
			sum := nums[i] + nums[i-1]
			if sum < minSum {
				minSum = sum
				minIdx = i - 1
			}
		}

		newNums := []int{}
		for i := 0; i < len(nums); i++ {
			if i == minIdx {
				newNums = append(newNums, minSum)
				count++
				i++
			} else {
				newNums = append(newNums, nums[i])
			}
		}

		nums = newNums
	}

	return count
}
