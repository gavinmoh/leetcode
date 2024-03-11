package main

func customSortString(order string, s string) string {
	result := ""
	hash := make(map[rune]int)

	for _, char := range s {
		hash[char]++
	}

	for _, char := range order {
		if _, ok := hash[char]; !ok {
			continue
		}

		for ; hash[char] > 0; hash[char]-- {
			result += string(char)
		}
	}

	for char, count := range hash {
		for ; count > 0; count-- {
			result += string(char)
		}
	}

	return result
}
