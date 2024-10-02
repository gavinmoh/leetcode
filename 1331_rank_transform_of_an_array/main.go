package main

import "sort"

func arrayRankTransform(arr []int) []int {
	hash := make(map[int]int)
	ranking := []int{}
	for _, num := range arr {
		if _, ok := hash[num]; !ok {
			hash[num] = -1
			ranking = append(ranking, num)
		}
	}

	sort.Ints(ranking)
	for i := 0; i < len(ranking); i++ {
		hash[ranking[i]] = i + 1
	}

	for i := 0; i < len(arr); i++ {
		arr[i] = hash[arr[i]]
	}

	return arr
}
