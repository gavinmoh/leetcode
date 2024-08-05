package main

import "sort"

type Item struct {
	Str   string
	First int
	Count int
}

func kthDistinct(arr []string, k int) string {
	count := make(map[string]*Item)
	for i, str := range arr {
		if _, ok := count[str]; !ok {
			count[str] = &Item{
				Str:   str,
				First: i,
				Count: 1,
			}
		} else {
			count[str].Count += 1
		}
	}

	distincts := []*Item{}
	for _, item := range count {
		if item.Count == 1 {
			distincts = append(distincts, item)
		}
	}

	sort.SliceStable(distincts, func(i, j int) bool {
		return distincts[i].First < distincts[j].First
	})

	if len(distincts) < k {
		return ""
	}

	return distincts[k-1].Str
}
