package main

func maximumOddBinaryNumber(s string) string {
	count := 0
	for _, char := range s {
		if char == '1' {
			count++
		}
	}

	result := ""
	for i := 1; i < count; i++ {
		result += "1"
	}
	for j := 0; j < len(s)-count; j++ {
		result += "0"
	}

	return result + "1"
}
