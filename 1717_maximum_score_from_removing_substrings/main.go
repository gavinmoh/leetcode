package main

func maximumGain(s string, x int, y int) int {
	points := 0
	high, low := "ab", "ba"
	if y > x {
		high, low = low, high
		x, y = y, x
	}

	s, highCount := removeSubstring(s, high)
	points += highCount * x

	s, lowCount := removeSubstring(s, low)
	points += lowCount * y

	return points
}

func removeSubstring(s string, target string) (string, int) {
	result, count := []byte{}, 0
	for i := 0; i < len(s); i++ {
		result = append(result, s[i])
		n := len(result)

		if n < 2 {
			continue
		}

		if string(result[n-2:n]) == target {
			result = result[:n-2]
			count++
		}
	}

	return string(result), count
}
