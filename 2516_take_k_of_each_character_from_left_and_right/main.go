package main

func takeCharacters(s string, k int) int {
	if k == 0 {
		return 0
	}

	freq := make([]int, 3)
	for _, char := range s {
		freq[char-'a']++
	}

	for _, count := range freq {
		if count < k {
			return -1
		}
	}

	window := []int{0, 0, 0}
	left, maxWindow := 0, 0
	for right := 0; right < len(s); right++ {
		window[s[right]-'a']++

		for left <= right && (freq[0]-window[0] < k || freq[1]-window[1] < k || freq[2]-window[2] < k) {
			window[s[left]-'a']--
			left++
		}

		maxWindow = max(maxWindow, right-left+1)
	}

	return len(s) - maxWindow
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
