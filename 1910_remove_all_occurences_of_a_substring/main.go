package main

func removeOccurrences(s string, part string) string {
	result, target := []rune{}, []rune(part)
	for _, c := range s {
		result = append(result, c)

		m, n := len(result), len(part)
		if m < n {
			continue
		}

		if equal(result[m-n:m], target, n) {
			result = result[:m-n]
		}
	}

	return string(result)
}

func equal(arr1 []rune, arr2 []rune, n int) bool {
	for i := 0; i < n; i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}

	return true
}
