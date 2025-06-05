package main

import "strings"

func smallestEquivalentString(s1 string, s2 string, baseStr string) string {
	parent := make([]rune, 26)
	for i := rune(0); i < rune(26); i++ {
		parent[i] = i
	}

	var find func(i rune) rune
	find = func(i rune) rune {
		if parent[i] == i {
			return i
		}

		return find(parent[i])
	}

	join := func(i, j rune) {
		parentI := find(i)
		parentJ := find(j)

		if parentI < parentJ {
			parent[parentJ] = parentI
		} else {
			parent[parentI] = parentJ
		}
	}

	n := len(s1)
	for i := 0; i < n; i++ {
		c1, c2 := rune(s1[i])-'a', rune(s2[i])-'a'
		join(c1, c2)
	}

	var result strings.Builder
	for _, c := range baseStr {
		repC := find(c - 'a')
		result.WriteRune('a' + repC)
	}

	return result.String()
}
