package main

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}

	prefix := ""
	hasCommonPrefix := true

	for i := 0; hasCommonPrefix; i++ {
		if len(strs[0]) <= i {
			break
		}

		testChar := strs[0][i]
		for _, str := range strs {
			if len(str) <= i {
				hasCommonPrefix = false
				break
			}

			char := str[i]

			if testChar != char {
				hasCommonPrefix = false
				break
			}

		}
		if hasCommonPrefix {
			prefix = prefix + string(testChar)
		}
	}

	return prefix
}
