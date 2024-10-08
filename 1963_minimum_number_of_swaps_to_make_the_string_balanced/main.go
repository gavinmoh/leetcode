package main

func minSwaps(s string) int {
	letters := []rune(s)
	open, close := 0, 0
	right := len(s) - 1
	count := 0
	for left, letter := range letters {
		if letter == '[' {
			open++
		} else {
			close++
		}
		if close > open {
			for right > left {
				if letters[right] == '[' {
					letters[right] = ']'
					close--
					open++
					count++
					break
				}
				right--
			}
		}
	}

	return count
}
