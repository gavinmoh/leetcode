package main

func countHomogenous(s string) int {
	modulo := 1000000007
	count := 0
	str := ""
	for i := 0; i < len(s); i++ {
		if len(str) == 0 || str[0] == s[i] {
			str += string(s[i])
		} else {
			n := len(str)
			total := n * (n + 1) / 2 // k choose 2 formula
			count += total
			str = string(s[i])
		}
	}

	// final substring
	n := len(str)
	total := n * (n + 1) / 2 // k choose 2 formula
	count += total

	if count > modulo {
		count = count % modulo
	}

	return count
}
