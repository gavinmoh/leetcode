package main

import "sort"

func sortVowels(s string) string {
	n := len(s)
	vowels := []int{}
	strMap := make(map[int]rune, n)
	boolMap := make(map[int]bool, n)

	for i, char := range s {
		strMap[i] = char
		if isVowel(char) {
			boolMap[i] = true
			vowels = append(vowels, int(char))
		}
	}

	if len(vowels) == 0 {
		return s
	}

	sort.Ints(vowels)

	permuted := ""
	index := 0
	for i := 0; i < n; i++ {
		if boolMap[i] {
			permuted = permuted + string(rune(vowels[index]))
			index++
		} else {
			permuted = permuted + string(strMap[i])
		}
	}

	return permuted
}

func isVowel(char rune) bool {
	switch char {
	case 'A', 'E', 'I', 'O', 'U', 'a', 'e', 'i', 'o', 'u':
		return true
	}

	return false
}
