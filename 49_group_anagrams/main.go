package main

import (
	"sort"
	"strconv"
)

func groupAnagrams(strs []string) [][]string {
	anagramMap := make(map[string][]string)

	for _, word := range strs {
		letterCount := make(map[rune]int)
		uniqueLetters := []rune{}
		for _, letter := range word {
			if _, ok := letterCount[letter]; !ok {
				uniqueLetters = append(uniqueLetters, letter)
			}
			letterCount[letter]++
		}
		sort.Slice(uniqueLetters, func(i, j int) bool {
			return uniqueLetters[i] < uniqueLetters[j]
		})
		anagramKey := ""
		for _, letter := range uniqueLetters {
			anagramKey += string(letter) + strconv.Itoa(letterCount[letter])
		}
		if _, ok := anagramMap[anagramKey]; ok {
			anagramMap[anagramKey] = append(anagramMap[anagramKey], word)
		} else {
			anagramMap[anagramKey] = []string{word}
		}
	}

	anagrams := [][]string{}
	for _, words := range anagramMap {
		anagrams = append(anagrams, words)
	}

	return anagrams
}
