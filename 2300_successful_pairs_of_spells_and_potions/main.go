package main

import (
	"math"
	"sort"
)

func successfulPairs(spells []int, potions []int, success int64) []int {
	n, m := len(spells), len(potions)
	sort.Ints(potions)
	result := make([]int, n)

	for i, spell := range spells {
		minStrength := int(math.Ceil(float64(success) / float64(spell)))
		minIdx := sort.SearchInts(potions, minStrength)
		result[i] = m - minIdx
	}

	return result
}
