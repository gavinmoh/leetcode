package main

func numTeams(rating []int) int {
	count := 0
	for i, num := range rating {
		left := rating[:i]
		right := rating[i+1:]

		leftLessThan, leftMoreThan := 0, 0
		for _, leftNum := range left {
			if leftNum < num {
				leftLessThan++
			} else if leftNum > num {
				leftMoreThan++
			}
		}

		rightLessThan, rightMoreThan := 0, 0
		for _, rightNum := range right {
			if rightNum < num {
				rightLessThan++
			} else if rightNum > num {
				rightMoreThan++
			}
		}

		count += leftLessThan * rightMoreThan
		count += leftMoreThan * rightLessThan
	}

	return count
}
