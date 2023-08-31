package main

func lengthOfLongestSubstring(s string) int {
	// var longestSubstring int
	// var currentSubstring int
	// var currentSubstringMap = make(map[string]int)
	substring := ""
	longestSubstring := ""

	for _, char := range s {
		if len(substring) == 0 {
			substring = string(char)
		} else if len(substring) == 1 {
			if substring != string(char) {
				substring = substring + string(char)
			}
		} else {
			var substringMap = make(map[rune]int)
			for i, c := range substring {
				substringMap[c] = i
			}
			if index, ok := substringMap[char]; ok {
				if substring[0] == byte(char) {
					substring = substring[1:] + string(char)
				} else {
					substring = substring[index+1:] + string(char)
				}
			} else {
				substring = substring + string(char)
			}
		}

		if len(substring) > len(longestSubstring) {
			longestSubstring = substring
		}
	}

	return len(longestSubstring)
}
