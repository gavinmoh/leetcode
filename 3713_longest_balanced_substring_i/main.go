package main

func longestBalanced(s string) int {
	maxLength := 0

	for left := 0; left < len(s); left++ {
		freq := make([]int, 26)

		for right := left; right < len(s); right++ {
			freq[s[right]-byte('a')]++

			balanced := true
			nonZero := -1
			for _, count := range freq {
				if count == 0 {
					continue
				}

				if nonZero == -1 {
					nonZero = count
				} else if nonZero != count {
					balanced = false
					break
				}
			}

			if balanced {
				length := right - left + 1
				if length > maxLength {
					maxLength = length
				}
			}
		}
	}

	return maxLength
}
