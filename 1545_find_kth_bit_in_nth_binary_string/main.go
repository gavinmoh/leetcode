package main

func findKthBit(n int, k int) byte {
	s := "0"
	for i := 1; i < n; i++ {
		s = s + "1" + reverseInvert(s)
	}

	return s[k-1]
}

func reverseInvert(s string) string {
	result := make([]byte, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '0' {
			result[len(s)-i-1] = '1'
		} else {
			result[len(s)-i-1] = '0'
		}
	}

	return string(result)
}
