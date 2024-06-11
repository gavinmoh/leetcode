package main

import "sort"

func relativeSortArray(arr1 []int, arr2 []int) []int {
	arr2Map := make(map[int]bool)
	for _, num := range arr2 {
		arr2Map[num] = true
	}

	extras := []int{}
	arr1Map := make(map[int]int)
	for _, num := range arr1 {
		if _, ok := arr2Map[num]; !ok {
			extras = append(extras, num)
		} else {
			arr1Map[num]++
		}
	}
	sort.Ints(extras)

	sorted := []int{}
	for _, num := range arr2 {
		if count, ok := arr1Map[num]; ok {
			for i := 0; i < count; i++ {
				sorted = append(sorted, num)
			}
		}
	}

	return append(sorted, extras...)
}
