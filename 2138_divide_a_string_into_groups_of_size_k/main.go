package main

import "strings"

func divideString(s string, k int, fill byte) []string {
	result := []string{}
	n := len(s)

	var builder strings.Builder
	for i := 0; i < n; i++ {
		builder.WriteByte(s[i])

		if builder.Len() == k {
			result = append(result, builder.String())
			builder.Reset()
		}
	}

	// fill in the rest with fill character
	if builder.Len() > 0 {
		remain := k - builder.Len()
		for i := 0; i < remain; i++ {
			builder.WriteByte(fill)
		}

		result = append(result, builder.String())
	}

	return result
}
