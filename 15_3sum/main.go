package main

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

func ThreeSum(nums []int) [][]int {
	result := [][]int{}

	nums = SortArray(nums)

	for i := 0; i < len(nums); i++ {
		firstNumber := nums[i]

		if i != 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < len(nums); j++ {
			secondNumber := nums[j]

			required := 0 - firstNumber - secondNumber

			for k := j + 1; k < len(nums); k++ {
				thirdNumber := nums[k]

				if thirdNumber > required {
					break
				}

				if !(i != j && i != k && j != k) {
					continue
				}

				if thirdNumber == required {
					triplet := []int{firstNumber, secondNumber, thirdNumber}
					// check if the result already exists
					exists := false
					for _, r := range result {
						if r[0] == triplet[0] && r[1] == triplet[1] && r[2] == triplet[2] {
							exists = true
							continue
						}
					}
					if exists {
						continue
					}

					result = append(result, triplet)
					break
				}
			}
		}
	}

	return result
}
