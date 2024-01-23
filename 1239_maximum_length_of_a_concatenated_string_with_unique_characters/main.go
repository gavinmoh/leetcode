package main

func maxLength(arr []string) int {
	n := len(arr)
	cache := make(map[string]bool)

	var dfs func(int, string) int
	dfs = func(i int, acc string) int {
		if i >= n {
			return len(acc)
		}

		newStr := acc + arr[i]
		if _, ok := cache[newStr]; !ok {
			cache[newStr] = isCharUnique(newStr)
		}

		// if it is unique
		if cache[newStr] {
			left := dfs(i+1, newStr)
			right := dfs(i+1, acc)

			return max(left, right)
		}

		return dfs(i+1, acc)
	}

	return dfs(0, "")
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func isCharUnique(str string) bool {
	charMap := make(map[rune]int)
	for _, char := range str {
		charMap[char]++
		if charMap[char] > 1 {
			return false
		}
	}
	return true
}
