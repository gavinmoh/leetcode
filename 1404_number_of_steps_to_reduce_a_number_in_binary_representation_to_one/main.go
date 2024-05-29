package main

func numSteps(s string) int {
	count := 0
	for s != "1" {
		lastChar := s[len(s)-1]
		if lastChar == '1' {
			s = stringAdd(s, 1)
		} else {
			s = s[:len(s)-1]
		}
		count++
	}

	return count
}

func stringAdd(s string, x int) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '1' && x > 0 {
			s = s[:i] + "0" + s[i+1:]
		} else if s[i] == '0' && x > 0 {
			s = s[:i] + "1" + s[i+1:]
			x--
		}
	}

	for x > 0 {
		s = "1" + s
		x--
	}

	return s
}
