package main

import "sort"

type Character struct {
	Character rune
	Count     int
}

func frequencySort(s string) string {
	counts := make(map[rune]int)

	for _, char := range s {
		counts[char]++
	}

	characters := []*Character{}
	for char, count := range counts {
		newCharacter := &Character{Character: char, Count: count}
		characters = append(characters, newCharacter)
	}

	sort.Slice(characters, func(i, j int) bool {
		return characters[i].Count > characters[j].Count
	})

	result := ""
	for _, character := range characters {
		for i := 0; i < character.Count; i++ {
			result += string(character.Character)
		}
	}

	return result
}
