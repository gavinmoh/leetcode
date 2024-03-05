package main

func minimumLength(s string) int {
	left, right := 0, len(s)-1
	for right > left {
		if s[left] == s[right] {
			left++
			right--
			for left <= right && s[left] == s[left-1] {
				left++
			}
			for right >= left && s[right] == s[right+1] {
				right--
			}
		} else {
			break
		}
	}
	return right - left + 1
}
