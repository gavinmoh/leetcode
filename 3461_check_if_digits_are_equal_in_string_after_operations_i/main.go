package main

import "strings"

func hasSameDigits(s string) bool {
	for len(s) > 2 {
		var newS strings.Builder
		for i := 0; i < len(s)-1; i++ {
			d1 := int(s[i] - '0')
			d2 := int(s[i+1] - '0')
			newDigit := (d1 + d2) % 10
			newS.WriteRune('0' + rune(newDigit))
		}

		s = newS.String()
	}

	return s[0] == s[1]
}
