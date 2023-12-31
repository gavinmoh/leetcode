package main

type Character struct {
	char  rune
	first int
	last  int
}

func (c *Character) diff() int {
	return c.last - c.first - 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxLengthBetweenEqualCharacters(s string) int {
	charMap := make(map[rune]*Character)
	maxDiff := -1

	for i, char := range s {
		if character, ok := charMap[char]; !ok {
			charMap[char] = &Character{char: char, first: i, last: i}
		} else {
			character.last = i
			maxDiff = max(maxDiff, character.diff())
		}
	}

	return maxDiff
}
