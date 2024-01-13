package main

func minSteps(s string, t string) int {
	sMap := make(map[byte]int)
	tMap := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		sMap[s[i]]++
		tMap[t[i]]++
	}

	step := 0
	for char, freq := range sMap {
		if tMap[char] < freq {
			step += freq - tMap[char]
		}
	}

	return step
}
