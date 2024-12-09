package main

func maximumLength(s string) int {
	frequency := make([][]int, 26)
	for i := 0; i < 26; i++ {
		frequency[i] = make([]int, len(s)+1)
	}
	substrLength := 1
	prev := s[0]
	frequency[int(prev-'a')][1] = 1

	for i := 1; i < len(s); i++ {
		char := s[i]

		if char == prev {
			substrLength++
			frequency[char-'a'][substrLength]++
		} else {
			prev = char
			substrLength = 1
			frequency[char-'a'][1]++
		}
	}

	maxLength := -1
	for char := 0; char < 26; char++ {
		for length := len(s) - 1; length >= 0; length-- {
			frequency[char][length] += frequency[char][length+1]
			if frequency[char][length] >= 3 {
				maxLength = max(maxLength, length)
				break
			}
		}
	}

	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
