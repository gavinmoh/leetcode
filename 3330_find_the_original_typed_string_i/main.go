package main

func possibleStringCount(word string) int {
	count := 1
	currentCount := 1
	currentChar := word[0]

	for i := 1; i < len(word); i++ {
		if currentChar != word[i] {
			count += currentCount - 1

			// reset
			currentCount = 1
			currentChar = word[i]
		} else {
			currentCount++
		}
	}
	count += currentCount - 1

	return count
}
