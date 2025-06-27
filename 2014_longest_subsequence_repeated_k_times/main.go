package main

func longestSubsequenceRepeatedK(s string, k int) string {
	freq := make([]int, 26)
	for _, c := range s {
		freq[int(c-'a')]++
	}

	usables := []rune{}
	for i := 25; i >= 0; i-- {
		if freq[i] < k {
			continue
		}

		c := rune(i) + 'a'
		usables = append(usables, c)
	}

	if len(usables) == 0 {
		return ""
	}

	isKRepeated := func(pattern string) bool {
		i, matched := 0, 0
		for _, c := range s {
			if c == rune(pattern[i]) {
				i++
				if i == len(pattern) {
					i = 0 // reset i to 0
					matched++
					if matched == k {
						return true
					}
				}
			}
		}

		return false
	}

	candidates := []string{}
	for _, c := range usables {
		candidates = append(candidates, string(c))
	}

	result := ""
	for len(candidates) > 0 {
		current := string(candidates[0])
		candidates = candidates[1:]
		if len(current) > len(result) {
			result = current
		}

		// generate next string
		for _, c := range usables {
			next := current + string(c)
			if isKRepeated(next) {
				candidates = append(candidates, next)
			}
		}
	}

	return result
}
