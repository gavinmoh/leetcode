package main

import "sort"

func findContentChildren(g []int, s []int) int {
	childs := len(g)
	content := 0

	sort.Ints(g)
	sort.Ints(s)

	childIndex := 0
	for _, cookieSize := range s {
		if childIndex >= childs {
			break
		}
		if cookieSize >= g[childIndex] {
			childIndex++
			content++
		}
	}

	return content
}
