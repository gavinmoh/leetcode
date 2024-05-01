package main

func reversePrefix(word string, ch byte) string {
	n := len(word)
	result := ""

	for i := 0; i < n; i++ {
		if word[i] != ch {
			continue
		}
		for j := i; j >= 0; j-- {
			result += string(word[j])
		}
		for k := i + 1; k < n; k++ {
			result += string(word[k])
		}
		break
	}

	if len(result) > 1 {
		return result
	}

	return word
}
