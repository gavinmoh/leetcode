package main

import "math"

func hasAllCodes(s string, k int) bool {
	substrings := map[string]bool{}

	for i := 0; i <= len(s)-k; i++ {
		substr := s[i : i+k]
		substrings[substr] = true
	}

	expected := int(math.Pow(float64(2), float64(k)))
	return expected == len(substrings)
}
