package main

func largestOddNumber(num string) string {
	lastIntegerIndex := -1

	for i, char := range num {
		if int(char)%2 != 0 {
			lastIntegerIndex = i
		}
	}

	if lastIntegerIndex != -1 {
		return num[0 : lastIntegerIndex+1]
	}

	return ""

}
