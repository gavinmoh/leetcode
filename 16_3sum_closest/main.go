package main

func SortArray(nums []int) []int {
	swapped := false
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-1; j++ {
			if nums[j] > nums[j+1] {
				swapped = true
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}

		if !swapped {
			break
		}
	}
	return nums
}

func ThreeSumClosest(nums []int, target int) int {
	nums = SortArray(nums)
	difference := 0
	sum := 0

	for i := 0; i < len(nums)-2; i++ {
		first := nums[i]

		for j := i + 1; j < len(nums)-1; j++ {
			second := nums[j]

			for k := j + 1; k < len(nums); k++ {
				third := nums[k]

				currentSum := first + second + third
				currentDifference := target - currentSum

				// make the difference positive
				if currentDifference < 0 {
					currentDifference = currentDifference * -1
				}

				// first loop, we initialize the difference
				// then we continue to next loop
				if i == 0 && j == i+1 && k == j+1 {
					difference = currentDifference
					sum = currentSum
					continue
				}

				// if the difference is 0, we found the target
				if currentDifference == 0 {
					return currentSum
				}

				// if the current difference is greater than the difference, we can skip
				if currentDifference >= difference {
					continue
				}

				// if the current difference is less than the difference, we update the difference
				difference = currentDifference
				sum = currentSum
			}
		}
	}

	return sum
}
