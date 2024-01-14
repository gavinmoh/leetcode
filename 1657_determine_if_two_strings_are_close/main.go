package main

import "sort"

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	word1Map := make(map[byte]int)
	word2Map := make(map[byte]int)

	for i := 0; i < len(word1); i++ {
		word1Map[word1[i]] += 1
		word2Map[word2[i]] += 1
	}

	word1Freq := []int{}
	word2Freq := []int{}

	for char, freq := range word1Map {
		if _, ok := word2Map[char]; !ok {
			return false
		}
		word1Freq = append(word1Freq, freq)
	}

	for char, freq := range word2Map {
		if _, ok := word1Map[char]; !ok {
			return false
		}
		word2Freq = append(word2Freq, freq)
	}

	sort.Ints(word1Freq)
	sort.Ints(word2Freq)

	if len(word1Freq) != len(word2Freq) {
		return false
	}

	for i := 0; i < len(word1Freq); i++ {
		if word1Freq[i] != word2Freq[i] {
			return false
		}
	}

	return true
}
