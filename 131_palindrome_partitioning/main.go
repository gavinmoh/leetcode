package main

func partition(s string) [][]string {
	palindromeCache := make(map[string]bool)
	isPalindrome := func(str string) bool {
		if cachedResult, ok := palindromeCache[str]; ok {
			return cachedResult
		}

		left, right := 0, len(str)-1
		for left < right {
			if str[left] != str[right] {
				palindromeCache[str] = false
				return false
			}
			left++
			right--
		}
		palindromeCache[str] = true
		return true
	}

	n := len(s)
	var dfs func(left int) [][]string
	dfs = func(left int) [][]string {
		if left > n-1 {
			return [][]string{{}}
		}

		output := [][]string{}
		for right := left + 1; right <= n; right++ {
			substr := s[left:right]
			if isPalindrome(substr) {
				rightPartitioned := dfs(right)
				for _, partitioned := range rightPartitioned {
					partitioned = append([]string{substr}, partitioned...)
					output = append(output, partitioned)
				}
			}
		}

		return output
	}

	return dfs(0)
}
