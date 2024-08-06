package main

import "sort"

type Item struct {
	Letter byte
	Freq   int
	Push   int
}

func minimumPushes(word string) int {
	count := make([]*Item, 26)
	for i := 0; i < 26; i++ {
		item := &Item{
			Letter: 'a' + byte(i),
			Freq:   0,
		}
		count[i] = item
	}

	for _, letter := range word {
		count[letter-'a'].Freq += 1
	}

	sort.SliceStable(count, func(i, j int) bool {
		return count[i].Freq > count[j].Freq
	})

	i := 0
	for j := 1; i < 26; j++ {
		for k := 0; k < 8 && i < 26; k++ {
			count[i].Push = j
			i++
		}
	}

	result := 0
	for _, item := range count {
		result += (item.Freq * item.Push)
	}

	return result
}
