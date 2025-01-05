package main

import "strings"

func shiftingLetters(s string, shifts [][]int) string {
	n := len(s)
	diffArray := make([]int, len(s))
	for _, shift := range shifts {
		start, end, direction := shift[0], shift[1], shift[2]
		if direction == 1 {
			diffArray[start]++
			if end+1 < n {
				diffArray[end+1]--
			}
		} else {
			diffArray[start]--
			if end+1 < n {
				diffArray[end+1]++
			}
		}
	}

	var result strings.Builder
	numOfShifts := 0
	for i, char := range s {
		numOfShifts = (numOfShifts + diffArray[i]) % 26
		if numOfShifts < 0 {
			numOfShifts += 26
		}
		newChar := 'a' + (char-'a'+rune(numOfShifts))%26
		result.WriteRune(newChar)
	}

	return result.String()
}
