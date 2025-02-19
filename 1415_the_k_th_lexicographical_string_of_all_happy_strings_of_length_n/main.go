package main

import "sort"

func getHappyString(n int, k int) string {
	happyStrings := []string{}
	str := []rune{}

	var dfs func(length int)
	dfs = func(length int) {
		if length == 0 {
			happyStrings = append(happyStrings, string(str))
			return
		}

		for _, char := range []rune{'a', 'b', 'c'} {
			if len(str) > 0 && str[len(str)-1] == char {
				continue
			}
			str = append(str, char)
			dfs(length - 1)
			str = str[:len(str)-1]
		}
	}

	dfs(n)

	if len(happyStrings) < k {
		return ""
	}

	sort.SliceStable(happyStrings, func(i, j int) bool {
		return happyStrings[i] < happyStrings[j]
	})

	return happyStrings[k-1]
}
