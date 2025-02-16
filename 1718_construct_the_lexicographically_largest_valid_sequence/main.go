package main

func constructDistancedSequence(n int) []int {
	result := make([]int, n*2-1)
	isNumberUsed := make([]bool, n+1)

	var findLexicographicallyLargestSequence func(idx int) bool
	findLexicographicallyLargestSequence = func(idx int) bool {
		if idx == len(result) {
			return true
		}

		if result[idx] != 0 {
			return findLexicographicallyLargestSequence(idx + 1)
		}

		for num := n; num > 0; num-- {
			if isNumberUsed[num] {
				continue
			}

			isNumberUsed[num] = true
			result[idx] = num

			if num == 1 {
				if findLexicographicallyLargestSequence(idx + 1) {
					return true
				}
			} else if (idx+num) < len(result) && result[idx+num] == 0 {
				result[idx+num] = num

				if findLexicographicallyLargestSequence(idx + 1) {
					return true
				}

				result[idx+num] = 0
			}

			result[idx] = 0
			isNumberUsed[num] = false
		}

		return false
	}

	findLexicographicallyLargestSequence(0)

	return result
}
