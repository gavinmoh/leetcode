package main

import (
	"fmt"
)

func SortArray(nums []int) []int {
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(nums)-1; i++ {
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
				swapped = true
			}
		}
	}

	return nums
}

func FourSum(nums []int, target int) [][]int {
	results := [][]int{}
	quadruplets := make(map[string][]int) // key: string, value: []int
	nums = SortArray(nums)
	maxValue := nums[len(nums)-1]

	maxValueReached := false
	for i := 0; i < len(nums)-3; i++ {
		first := nums[i]

		// handling case where the rest of the numbers are all the same
		if maxValueReached {
			break
		} else if !maxValueReached && first == maxValue {
			maxValueReached = true
		}

		for j := i + 1; j < len(nums)-2; j++ {
			second := nums[j]

			for k := j + 1; k < len(nums)-1; k++ {
				third := nums[k]

				for l := k + 1; l < len(nums); l++ {
					fourth := nums[l]

					sum := first + second + third + fourth
					if sum == target {
						combo := []int{first, second, third, fourth}
						comboKey := fmt.Sprintf("%d;%d;%d;%d;", first, second, third, fourth)
						// check if comboKey exists in quadruplets
						if _, ok := quadruplets[comboKey]; !ok {
							quadruplets[comboKey] = combo
							results = append(results, combo)
						}
					} else if sum > target {
						break
					}
				}
			}
		}
	}

	return results
}
